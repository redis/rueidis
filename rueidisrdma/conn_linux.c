/*
 * Copyright (c) 2024, zhenwei pi <pizhenwei@bytedance.com>
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *   * Redistributions of source code must retain the above copyright notice,
 *     this list of conditions and the following disclaimer.
 *   * Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *   * Neither the name of the copyright holder nor the names of its
 *     contributors may be used to endorse or promote products derived
 *     from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

// Modified from https://github.com/valkey-io/valkey/blob/unstable/deps/libvalkey/src/rdma.c

#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <fcntl.h>
#include <poll.h>
#include <assert.h>
#include <sys/socket.h>
#include <sys/time.h>
#include "conn_linux.h"

/* Return the current time in microseconds since Epoch */
static inline int64_t vk_usec_now(void) {
    int64_t usec;
    struct timeval now;
    if (gettimeofday(&now, NULL) < 0) {
        return -1;
    }
    usec = (int64_t)now.tv_sec * 1000000LL + (int64_t)now.tv_usec;
    return usec;
}

static inline int64_t vk_msec_now(void) {
    return vk_usec_now() / 1000;
}

static inline int valkeyMin(long long a, long long b) {
    return (a < b) ? a : b;
}

static inline int poll_noeintr(struct pollfd *fds, nfds_t nfds, int timeout_ms) {
    for (;;) {
        int rc = poll(fds, nfds, timeout_ms);
        if (rc >= 0) return rc;
        if (errno == EINTR) continue;
        return -1;
    }
}

void valkeySetError(RdmaContext *ctx, int type, const char *str) {
    size_t len;
    len = strlen(str);
    pthread_mutex_lock(&ctx->err_mu);
    len = len < (sizeof(ctx->errstr) - 1) ? len : (sizeof(ctx->errstr) - 1);
    ctx->err = type;
    memcpy(ctx->errstr, str, len);
    ctx->errstr[len] = '\0';
    pthread_mutex_unlock(&ctx->err_mu);
}

static int valkeyRdmaCM(RdmaContext *ctx, long timeout);
static int connRdmaHandleCq(RdmaContext *ctx);

static int valkeyRdmaSetFdBlocking(RdmaContext *ctx, int fd, int blocking) {
    int flags;

    if ((flags = fcntl(fd, F_GETFL)) == -1) {
        valkeySetError(ctx, VALKEY_ERR_IO, "fcntl(F_GETFL)");
        return VALKEY_ERR;
    }

    if (blocking)
        flags &= ~O_NONBLOCK;
    else
        flags |= O_NONBLOCK;

    if (fcntl(fd, F_SETFL, flags) == -1) {
        valkeySetError(ctx, VALKEY_ERR_IO, "fcntl(F_SETFL)");
        return VALKEY_ERR;
    }

    return 0;
}

static int rdmaPostRecv(RdmaContext *ctx, struct rdma_cm_id *cm_id, valkeyRdmaCmd *cmd) {
    struct ibv_sge sge;
    size_t length = sizeof(valkeyRdmaCmd);
    struct ibv_recv_wr recv_wr, *bad_wr;

    sge.addr = (uint64_t)(uintptr_t)cmd;
    sge.length = length;
    sge.lkey = ctx->cmd_mr->lkey;

    recv_wr.wr_id = (uint64_t)cmd;
    recv_wr.sg_list = &sge;
    recv_wr.num_sge = 1;
    recv_wr.next = NULL;

    if (ibv_post_recv(cm_id->qp, &recv_wr, &bad_wr)) {
        return VALKEY_ERR;
    }

    return VALKEY_OK;
}

static void rdmaDestroyIoBuf(RdmaContext *ctx) {
    if (ctx->recv_mr) {
        ibv_dereg_mr(ctx->recv_mr);
        ctx->recv_mr = NULL;
    }

    free(ctx->recv_buf);
    ctx->recv_buf = NULL;

    if (ctx->send_mr) {
        ibv_dereg_mr(ctx->send_mr);
        ctx->send_mr = NULL;
    }

    free(ctx->send_buf);
    ctx->send_buf = NULL;

    if (ctx->cmd_mr) {
        ibv_dereg_mr(ctx->cmd_mr);
        ctx->cmd_mr = NULL;
    }

    free(ctx->cmd_buf);
    ctx->cmd_buf = NULL;
}

static int rdmaSetupIoBuf(RdmaContext *ctx, struct rdma_cm_id *cm_id) {
    int access = IBV_ACCESS_LOCAL_WRITE;
    size_t length = sizeof(valkeyRdmaCmd) * VALKEY_RDMA_MAX_WQE * 2;
    valkeyRdmaCmd *cmd;
    int i;

    /* setup CMD buf & MR */
    ctx->cmd_buf = calloc(length, 1);
    ctx->cmd_mr = ibv_reg_mr(ctx->pd, ctx->cmd_buf, length, access);
    if (!ctx->cmd_mr) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: reg cmd mr failed");
        goto destroy_iobuf;
    }

    for (i = 0; i < VALKEY_RDMA_MAX_WQE; i++) {
        cmd = ctx->cmd_buf + i;

        if (rdmaPostRecv(ctx, cm_id, cmd) == VALKEY_ERR) {
            valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: post recv failed");
            goto destroy_iobuf;
        }
    }

    for (i = VALKEY_RDMA_MAX_WQE; i < VALKEY_RDMA_MAX_WQE * 2; i++) {
        cmd = ctx->cmd_buf + i;
        cmd->keepalive.opcode = VALKEY_RDMA_INVALID_OPCODE;
    }

    /* setup recv buf & MR */
    access = IBV_ACCESS_LOCAL_WRITE | IBV_ACCESS_REMOTE_READ | IBV_ACCESS_REMOTE_WRITE;
    length = VALKEY_RDMA_DEFAULT_RX_LEN;
    ctx->recv_buf = calloc(length, 1);
    if (!ctx->recv_buf) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: receive buffer allocation failed");
        goto destroy_iobuf;
    }
    ctx->recv_length = length;
    ctx->recv_mr = ibv_reg_mr(ctx->pd, ctx->recv_buf, length, access);
    if (!ctx->recv_mr) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: reg recv mr failed");
        goto destroy_iobuf;
    }

    return VALKEY_OK;

destroy_iobuf:
    rdmaDestroyIoBuf(ctx);
    return VALKEY_ERR;
}

static int rdmaAdjustSendbuf(RdmaContext *ctx, unsigned int length) {
    int access = IBV_ACCESS_LOCAL_WRITE | IBV_ACCESS_REMOTE_READ | IBV_ACCESS_REMOTE_WRITE;

    if (length == ctx->send_length) {
        return VALKEY_OK;
    }

    /* try to free old MR & buffer */
    if (ctx->send_length) {
        ibv_dereg_mr(ctx->send_mr);
        free(ctx->send_buf);
        ctx->send_length = 0;
    }

    /* create a new buffer & MR */
    ctx->send_buf = calloc(length, 1);
    if (!ctx->send_buf) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: send buffer allocation failed");
        return VALKEY_ERR;
    }
    ctx->send_length = length;
    ctx->send_mr = ibv_reg_mr(ctx->pd, ctx->send_buf, length, access);
    if (!ctx->send_mr) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: reg send buf mr failed");
        free(ctx->send_buf);
        ctx->send_buf = NULL;
        ctx->send_length = 0;
        return VALKEY_ERR;
    }

    return VALKEY_OK;
}

static int rdmaSendCommand(RdmaContext *ctx, struct rdma_cm_id *cm_id, valkeyRdmaCmd *cmd) {
    struct ibv_send_wr send_wr, *bad_wr;
    struct ibv_sge sge;
    valkeyRdmaCmd *_cmd;
    int i;
    int ret;

find:
    /* find an unused cmd buffer */
    for (i = VALKEY_RDMA_MAX_WQE; i < 2 * VALKEY_RDMA_MAX_WQE; i++) {
        _cmd = ctx->cmd_buf + i;
        if (_cmd->keepalive.opcode == VALKEY_RDMA_INVALID_OPCODE) {
            _cmd->keepalive.opcode = 0; // claim
            break;
        }
    }

    if (i >= 2 * VALKEY_RDMA_MAX_WQE) {
        pthread_mutex_unlock(&ctx->rx_mu);
        if (connRdmaHandleCq(ctx) == VALKEY_ERR) {
            pthread_mutex_lock(&ctx->rx_mu);
            valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: failed to release cmd buf");
            return VALKEY_ERR;
        }
        pthread_mutex_lock(&ctx->rx_mu);
        goto find;
    }

    memcpy(_cmd, cmd, sizeof(valkeyRdmaCmd));
    sge.addr = (uint64_t)(uintptr_t)_cmd;
    sge.length = sizeof(valkeyRdmaCmd);
    sge.lkey = ctx->cmd_mr->lkey;

    send_wr.sg_list = &sge;
    send_wr.num_sge = 1;
    send_wr.wr_id = (uint64_t)_cmd;
    send_wr.opcode = IBV_WR_SEND;
    send_wr.send_flags = IBV_SEND_SIGNALED;
    send_wr.next = NULL;

resend:
    ret = ibv_post_send(cm_id->qp, &send_wr, &bad_wr);
    if (ret) {
        if (ret == ENOMEM) {
            pthread_mutex_unlock(&ctx->rx_mu);
            if (connRdmaHandleCq(ctx) == VALKEY_ERR) {
                pthread_mutex_lock(&ctx->rx_mu);
                valkeySetError(ctx, ret, "RDMA: failed to handle rx ENOMEM");
                return VALKEY_ERR;
            }
            pthread_mutex_lock(&ctx->rx_mu);
            goto resend;
        }
        valkeySetError(ctx, ret, "RDMA: failed to send command buffers");
        return VALKEY_ERR;
    }

    return VALKEY_OK;
}

static int connRdmaRegisterRx(RdmaContext *ctx, struct rdma_cm_id *cm_id) {
    valkeyRdmaCmd cmd = {0};

    cmd.memory.opcode = htons(RegisterXferMemory);
    cmd.memory.addr = htobe64((uint64_t)ctx->recv_buf);
    cmd.memory.length = htonl(ctx->recv_length);
    cmd.memory.key = htonl(ctx->recv_mr->rkey);

    ctx->rx_offset = 0;
    ctx->recv_offset = 0;

    return rdmaSendCommand(ctx, cm_id, &cmd);
}

static int connRdmaHandleRecv(RdmaContext *ctx, struct rdma_cm_id *cm_id, valkeyRdmaCmd *cmd, uint32_t byte_len) {
    int ret;
    if (byte_len != sizeof(valkeyRdmaCmd)) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: FATAL error, recv corrupted cmd");
        return VALKEY_ERR;
    }

    switch (ntohs(cmd->keepalive.opcode)) {
    case RegisterXferMemory:
        pthread_mutex_lock(&ctx->tx_mu);
        ctx->tx_addr = (char *)be64toh(cmd->memory.addr);
        ctx->tx_length = ntohl(cmd->memory.length);
        ctx->tx_key = ntohl(cmd->memory.key);
        ctx->tx_offset = 0;
        ret = rdmaAdjustSendbuf(ctx, ctx->tx_length);
        pthread_mutex_unlock(&ctx->tx_mu);
        if (ret == VALKEY_ERR) {
            return VALKEY_ERR;
        }
        break;

    case Keepalive:
        break;

    default:
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: FATAL error, unknown cmd");
        return VALKEY_ERR;
    }

    return rdmaPostRecv(ctx, cm_id, cmd);
}

static int connRdmaHandleRecvImm(RdmaContext *ctx, struct rdma_cm_id *cm_id, valkeyRdmaCmd *cmd, uint32_t byte_len) {
    pthread_mutex_lock(&ctx->rx_mu);
    assert(byte_len + ctx->rx_offset <= ctx->recv_length);
    ctx->rx_offset += byte_len;
    pthread_mutex_unlock(&ctx->rx_mu);

    return rdmaPostRecv(ctx, cm_id, cmd);
}

static int connRdmaHandleSend(RdmaContext *ctx, valkeyRdmaCmd *cmd) {
    /* mark this cmd has already sent */
    pthread_mutex_lock(&ctx->rx_mu);
    memset(cmd, 0x00, sizeof(*cmd));
    cmd->keepalive.opcode = VALKEY_RDMA_INVALID_OPCODE;
    pthread_mutex_unlock(&ctx->rx_mu);

    return VALKEY_OK;
}

static int connRdmaHandleWrite(RdmaContext *ctx, uint32_t byte_len) {
    return VALKEY_OK;
}

static int connRdmaHandleWc(RdmaContext *ctx, struct ibv_wc *wc) {
    struct rdma_cm_id *cm_id = ctx->cm_id;
    valkeyRdmaCmd *cmd;

    if (wc->status != IBV_WC_SUCCESS) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: send/recv failed");
        return VALKEY_ERR;
    }

    switch (wc->opcode) {
    case IBV_WC_RECV:
        cmd = (valkeyRdmaCmd *)(uintptr_t)wc->wr_id;
        if (connRdmaHandleRecv(ctx, cm_id, cmd, wc->byte_len) == VALKEY_ERR) {
            return VALKEY_ERR;
        }
        break;
    case IBV_WC_RECV_RDMA_WITH_IMM:
        cmd = (valkeyRdmaCmd *)(uintptr_t)wc->wr_id;
        if (connRdmaHandleRecvImm(ctx, cm_id, cmd, ntohl(wc->imm_data)) == VALKEY_ERR) {
            return VALKEY_ERR;
        }
        break;
    case IBV_WC_RDMA_WRITE:
        if (connRdmaHandleWrite(ctx, wc->byte_len) == VALKEY_ERR) {
            return VALKEY_ERR;
        }
        break;
    case IBV_WC_SEND:
        cmd = (valkeyRdmaCmd *)(uintptr_t)wc->wr_id;
        if (connRdmaHandleSend(ctx, cmd) == VALKEY_ERR) {
            return VALKEY_ERR;
        }
        break;
    default:
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: unexpected opcode");
        return VALKEY_ERR;
    }

    return VALKEY_OK;
}

static int connRdmaHandleCq(RdmaContext *ctx) {
    struct ibv_cq *ev_cq = NULL;
    void *ev_ctx = NULL;
    struct ibv_wc wc = {0};
    int ret;

    pthread_mutex_lock(&ctx->cq_mu);
    for (;;) {
        ret = ibv_poll_cq(ctx->cq, 1, &wc);
        if (ret < 0) {
            valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: poll cq failed");
            pthread_mutex_unlock(&ctx->cq_mu);
            return VALKEY_ERR;
        } else if (ret == 0) {
            break;
        }

        if (connRdmaHandleWc(ctx, &wc) == VALKEY_ERR) {
            pthread_mutex_unlock(&ctx->cq_mu);
            return VALKEY_ERR;
        }
    }

    if (ibv_get_cq_event(ctx->comp_channel, &ev_cq, &ev_ctx) < 0) {
        if (errno != EAGAIN) {
            valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: get cq event failed");
            pthread_mutex_unlock(&ctx->cq_mu);
            return VALKEY_ERR;
        }
        pthread_mutex_unlock(&ctx->cq_mu);
        return VALKEY_OK;
    }

    ibv_ack_cq_events(ev_cq, 1);
    if (ibv_req_notify_cq(ev_cq, 0)) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: notify cq failed");
        pthread_mutex_unlock(&ctx->cq_mu);
        return VALKEY_ERR;
    }

    for (;;) {
        ret = ibv_poll_cq(ctx->cq, 1, &wc);
        if (ret < 0) {
            valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: poll cq failed");
            pthread_mutex_unlock(&ctx->cq_mu);
            return VALKEY_ERR;
        } else if (ret == 0) {
            pthread_mutex_unlock(&ctx->cq_mu);
            return VALKEY_OK;
        }

        if (connRdmaHandleWc(ctx, &wc) == VALKEY_ERR) {
            pthread_mutex_unlock(&ctx->cq_mu);
            return VALKEY_ERR;
        }
    }
}

/* There are two FD(s) in use:
 * - fd of CM channel: handle CM event. Return error on Disconnected.
 * - fd of completion channel: handle CQ event.
 * Return OK on CQ event ready, then CQ event should be handled outside.
 */
static int valkeyRdmaPollCqCm(RdmaContext *ctx, long timed) {
#define VALKEY_RDMA_POLLFD_CM 0
#define VALKEY_RDMA_POLLFD_CQ 1
#define VALKEY_RDMA_POLLFD_MAX 2
    struct pollfd pfd[VALKEY_RDMA_POLLFD_MAX];
    struct ibv_wc wc = {0};
    long now;
    int ret;

    /* pfd[0] for CM event */
    pfd[VALKEY_RDMA_POLLFD_CM].fd = ctx->cm_channel->fd;
    pfd[VALKEY_RDMA_POLLFD_CM].events = POLLIN;
    pfd[VALKEY_RDMA_POLLFD_CM].revents = 0;

    /* pfd[1] for CQ event */
    pfd[VALKEY_RDMA_POLLFD_CQ].fd = ctx->comp_channel->fd;
    pfd[VALKEY_RDMA_POLLFD_CQ].events = POLLIN;
    pfd[VALKEY_RDMA_POLLFD_CQ].revents = 0;

    for (;;) {
        now = vk_msec_now();
        if (now >= timed) {
            valkeySetError(ctx, VALKEY_ERR_IO, "RDMA: IO timeout");
            return VALKEY_ERR;
        }

        /* First, try to drain CQ without relying on events. */
        for (;;) {
            pthread_mutex_lock(&ctx->cq_mu);
            ret = ibv_poll_cq(ctx->cq, 1, &wc);
            if (ret < 0) {
                valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: poll cq failed");
                pthread_mutex_unlock(&ctx->cq_mu);
                return VALKEY_ERR;
            } else if (ret == 0) {
                pthread_mutex_unlock(&ctx->cq_mu);
                break;
            }
            if (connRdmaHandleWc(ctx, &wc) == VALKEY_ERR) {
                pthread_mutex_unlock(&ctx->cq_mu);
                return VALKEY_ERR;
            }
            pthread_mutex_unlock(&ctx->cq_mu);
        }

        /* Poll for a short slice so we can re-check CQ even if no events. */
        ret = poll_noeintr(pfd, VALKEY_RDMA_POLLFD_MAX, (int)valkeyMin(10, timed - now));
        if (ret < 0) {
            valkeySetError(ctx, VALKEY_ERR_IO, "RDMA: Poll CQ/CM failed");
            return VALKEY_ERR;
        } else if (ret == 0) {
            continue;
        }

        if (pfd[VALKEY_RDMA_POLLFD_CM].revents & POLLIN) {
            valkeyRdmaCM(ctx, 0);
            if (!(ctx->flags & VALKEY_CONNECTED)) {
                valkeySetError(ctx, VALKEY_ERR_EOF, "Server closed the connection");
                return VALKEY_ERR;
            }
        }

        if (pfd[VALKEY_RDMA_POLLFD_CQ].revents & POLLIN) {
            if (connRdmaHandleCq(ctx) == VALKEY_ERR) {
                return VALKEY_ERR;
            }
            return VALKEY_OK;
        }
    }
}

static size_t connRdmaSend(RdmaContext *ctx, struct rdma_cm_id *cm_id, const void *data, size_t data_len) {
    struct ibv_send_wr send_wr, *bad_wr;
    struct ibv_sge sge;
    uint32_t off = ctx->tx_offset;
    char *addr = ctx->send_buf + off;
    char *remote_addr = ctx->tx_addr + off;
    int ret;

    assert(data_len <= ctx->tx_length);
    memcpy(addr, data, data_len);

    sge.addr = (uint64_t)(uintptr_t)addr;
    sge.lkey = ctx->send_mr->lkey;
    sge.length = data_len;

    send_wr.sg_list = &sge;
    send_wr.num_sge = 1;
    send_wr.opcode = IBV_WR_RDMA_WRITE_WITH_IMM;
    send_wr.send_flags = (++ctx->send_ops % VALKEY_RDMA_MAX_WQE) ? 0 : IBV_SEND_SIGNALED;
    send_wr.imm_data = htonl(data_len);
    send_wr.wr.rdma.remote_addr = (uint64_t)(uintptr_t)remote_addr;
    send_wr.wr.rdma.rkey = ctx->tx_key;
    send_wr.next = NULL;

resend:
    ret = ibv_post_send(cm_id->qp, &send_wr, &bad_wr);
    if (ret) {
        if (ret == ENOMEM) {
            pthread_mutex_unlock(&ctx->tx_mu);
            if (connRdmaHandleCq(ctx) == VALKEY_ERR) {
                pthread_mutex_lock(&ctx->tx_mu);
                valkeySetError(ctx, ret, "RDMA: failed to handle tx ENOMEM");
                return VALKEY_ERR;
            }
            pthread_mutex_lock(&ctx->tx_mu);
            goto resend;
        }
        valkeySetError(ctx, ret, "RDMA: failed to write with imm");
        return VALKEY_ERR;
    }

    ctx->tx_offset += data_len;

    return data_len;
}

static int valkeyRdmaConnect(RdmaContext *ctx, struct rdma_cm_id *cm_id) {
    struct ibv_comp_channel *comp_channel = NULL;
    struct ibv_cq *cq = NULL;
    struct ibv_pd *pd = NULL;
    struct ibv_qp_init_attr init_attr = {0};
    struct rdma_conn_param conn_param = {0};

    pd = ibv_alloc_pd(cm_id->verbs);
    if (!pd) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: alloc pd failed");
        goto error;
    }

    comp_channel = ibv_create_comp_channel(cm_id->verbs);
    if (!comp_channel) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: alloc comp channel failed");
        goto error;
    }

    if (valkeyRdmaSetFdBlocking(ctx, comp_channel->fd, 0) != VALKEY_OK) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: set recv comp channel fd non-block failed");
        goto error;
    }

    cq = ibv_create_cq(cm_id->verbs, VALKEY_RDMA_MAX_WQE * 2, ctx, comp_channel, 0);
    if (!cq) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: create send cq failed");
        goto error;
    }

    if (ibv_req_notify_cq(cq, 0)) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: notify send cq failed");
        goto error;
    }

    /* create qp with attr */
    init_attr.cap.max_send_wr = VALKEY_RDMA_MAX_WQE;
    init_attr.cap.max_recv_wr = VALKEY_RDMA_MAX_WQE;
    init_attr.cap.max_send_sge = 1;
    init_attr.cap.max_recv_sge = 1;
    init_attr.qp_type = IBV_QPT_RC;
    init_attr.send_cq = cq;
    init_attr.recv_cq = cq;
    if (rdma_create_qp(cm_id, pd, &init_attr)) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: create qp failed");
        goto error;
    }

    ctx->cm_id = cm_id;
    ctx->comp_channel = comp_channel;
    ctx->cq = cq;
    ctx->pd = pd;

    if (rdmaSetupIoBuf(ctx, cm_id) != VALKEY_OK)
        goto free_qp;

    /* rdma connect with param */
    conn_param.responder_resources = 1;
    conn_param.initiator_depth = 1;
    conn_param.retry_count = 7;
    conn_param.rnr_retry_count = 7;
    if (rdma_connect(cm_id, &conn_param)) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: connect failed");
        goto destroy_iobuf;
    }

    return VALKEY_OK;

destroy_iobuf:
    rdmaDestroyIoBuf(ctx);
free_qp:
    ibv_destroy_qp(cm_id->qp);
error:
    if (cq)
        ibv_destroy_cq(cq);
    if (pd)
        ibv_dealloc_pd(pd);
    if (comp_channel)
        ibv_destroy_comp_channel(comp_channel);

    return VALKEY_ERR;
}

static int valkeyRdmaEstablished(RdmaContext *ctx, struct rdma_cm_id *cm_id) {
    /* it's time to tell redis we have already connected */
    int ret;
    ctx->flags |= VALKEY_CONNECTED;
    pthread_mutex_lock(&ctx->rx_mu);
    ret = connRdmaRegisterRx(ctx, cm_id);
    pthread_mutex_unlock(&ctx->rx_mu);
    return ret;
}

static int valkeyRdmaCM(RdmaContext *ctx, long timeout) {
    struct rdma_cm_event *event;
    int ret = VALKEY_ERR;

    while (rdma_get_cm_event(ctx->cm_channel, &event) == 0) {
        switch (event->event) {
        case RDMA_CM_EVENT_ADDR_RESOLVED:
            if (timeout < 0 || timeout > 100)
                timeout = 100; /* at most 100ms to resolve route */
            ret = rdma_resolve_route(event->id, timeout);
            if (ret) {
                valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: route resolve failed on");
                goto disconnect;
            }
            break;
        case RDMA_CM_EVENT_ROUTE_RESOLVED:
            ret = valkeyRdmaConnect(ctx, event->id);
            if (ret == VALKEY_ERR) {
                goto disconnect;
            }
            break;
        case RDMA_CM_EVENT_ESTABLISHED:
            ret = valkeyRdmaEstablished(ctx, event->id);
            if (ret == VALKEY_ERR) {
                goto disconnect;
            }
            break;
        case RDMA_CM_EVENT_TIMEWAIT_EXIT:
        case RDMA_CM_EVENT_ADDR_ERROR:
        case RDMA_CM_EVENT_ROUTE_ERROR:
        case RDMA_CM_EVENT_CONNECT_ERROR:
        case RDMA_CM_EVENT_UNREACHABLE:
        case RDMA_CM_EVENT_REJECTED:
        case RDMA_CM_EVENT_DISCONNECTED:
            valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: disconnected");
            goto disconnect;
        case RDMA_CM_EVENT_ADDR_CHANGE:
        default:
            valkeySetError(ctx, VALKEY_ERR_OTHER, rdma_event_str(event->event));
            ret = VALKEY_ERR;
            break;
        }
        rdma_ack_cm_event(event);
    }

    return ret;

disconnect:
    rdma_ack_cm_event(event);
    ctx->flags &= ~VALKEY_CONNECTED;
    return VALKEY_ERR;
}

static int valkeyRdmaWaitTxBuf(RdmaContext *ctx, long timeout) {
    long now, end;

    assert(timeout >= 0);
    end = vk_msec_now() + timeout;

    while (1) {
        now = vk_msec_now();
        if (now >= end) {
            break;
        }
        if (valkeyRdmaPollCqCm(ctx, end) == VALKEY_OK) {
            if (connRdmaHandleCq(ctx) == VALKEY_ERR) {
                return VALKEY_ERR;
            }
            if (ctx->tx_length != 0 && ctx->send_length != 0) {
                return VALKEY_OK;
            }
        }
    }

    return VALKEY_ERR;
}

static int valkeyRdmaWaitConn(RdmaContext *ctx, long timeout) {
    struct pollfd pfd;
    long now, end;

    assert(timeout >= 0);
    end = vk_msec_now() + timeout;

    while (1) {
        now = vk_msec_now();
        if (now >= end) {
            break;
        }

        pfd.fd = ctx->cm_channel->fd;
        pfd.events = POLLIN;
        pfd.revents = 0;
        if (poll_noeintr(&pfd, 1, end - now) < 0) {
            valkeySetError(ctx, VALKEY_ERR_IO, "RDMA: Poll CM failed");
            return VALKEY_ERR;
        }

        if (valkeyRdmaCM(ctx, end - now) == VALKEY_ERR) {
            return VALKEY_ERR;
        }

        now = vk_msec_now();
        if (ctx->flags & VALKEY_CONNECTED) {
            if (valkeyRdmaWaitTxBuf(ctx, end - now) == VALKEY_OK) {
                return VALKEY_OK;
            }
        }
    }
    valkeySetError(ctx, VALKEY_ERR_TIMEOUT, "RDMA: connecting timeout");
    return VALKEY_ERR;
}

int rdmaConnect(RdmaContext *ctx, const char *addr, int port, long timeout_msec) {
    int ret;
    char _port[6]; /* strlen("65535"); */
    struct rdma_addrinfo hints = {0}, *addrinfo = NULL;
    long start = vk_msec_now(), timed;

    pthread_mutex_init(&ctx->cq_mu, NULL);
    pthread_mutex_init(&ctx->rx_mu, NULL);
    pthread_mutex_init(&ctx->tx_mu, NULL);
    pthread_mutex_init(&ctx->err_mu, NULL);

    ctx->tx_length = 0;
    ctx->send_length = 0;
    ctx->flags &= ~VALKEY_CONNECTED;

    ctx->cm_channel = rdma_create_event_channel();
    if (!ctx->cm_channel) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: create event channel failed");
        goto error;
    }

    if (rdma_create_id(ctx->cm_channel, &ctx->cm_id, (void *)ctx, RDMA_PS_TCP)) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: create id failed");
        goto error;
    }

    if ((valkeyRdmaSetFdBlocking(ctx, ctx->cm_channel->fd, 0) != VALKEY_OK)) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: set cm channel fd non-block failed");
        goto error;
    }

    /* resolve remote address & port by RDMA style */
    snprintf(_port, sizeof(_port), "%d", port);
    hints.ai_port_space = RDMA_PS_TCP;
    if (rdma_getaddrinfo(addr, _port, &hints, &addrinfo)) {
        valkeySetError(ctx, VALKEY_ERR_PROTOCOL, "RDMA: failed to getaddrinfo");
        goto error;
    }

    timed = timeout_msec - (vk_msec_now() - start);
    if (rdma_resolve_addr(ctx->cm_id, NULL, (struct sockaddr *)addrinfo->ai_dst_addr, timed)) {
        valkeySetError(ctx, VALKEY_ERR_OTHER, "RDMA: failed to resolve");
        goto error;
    }

    timed = vk_msec_now() - start;
    if (timed >= timeout_msec) {
        valkeySetError(ctx, VALKEY_ERR_TIMEOUT, "RDMA: resolving timeout");
        goto error;
    }

    if ((valkeyRdmaWaitConn(ctx, timeout_msec - timed) == VALKEY_OK) && (ctx->flags & VALKEY_CONNECTED)) {
        ret = VALKEY_OK;
        goto end;
    }

error:
    ret = VALKEY_ERR;
    if (ctx->cm_id) {
        rdma_destroy_id(ctx->cm_id);
    }
    if (ctx->cm_channel) {
        rdma_destroy_event_channel(ctx->cm_channel);
    }
    pthread_mutex_destroy(&ctx->cq_mu);
    pthread_mutex_destroy(&ctx->rx_mu);
    pthread_mutex_destroy(&ctx->tx_mu);
    pthread_mutex_destroy(&ctx->err_mu);
end:
    if (addrinfo) {
        rdma_freeaddrinfo(addrinfo);
    }
    return ret;
}

ssize_t rdmaRead(RdmaContext *ctx, char *buf, size_t bufcap, long timeout_msec) {
    struct rdma_cm_id *cm_id = ctx->cm_id;
    long end;
    uint32_t toread, remained, topoll;
    end = vk_msec_now() + timeout_msec;

pollcq:
    pthread_mutex_lock(&ctx->rx_mu);
    if (ctx->recv_offset < ctx->rx_offset) {
        remained = ctx->rx_offset - ctx->recv_offset;
        toread = valkeyMin(remained, bufcap);

        memcpy(buf, ctx->recv_buf + ctx->recv_offset, toread);
        ctx->recv_offset += toread;

        if (ctx->recv_offset == ctx->recv_length && connRdmaRegisterRx(ctx, cm_id) == VALKEY_ERR) {
            pthread_mutex_unlock(&ctx->rx_mu);
            return VALKEY_ERR;
        }
        pthread_mutex_unlock(&ctx->rx_mu);
        return toread;
    }
    pthread_mutex_unlock(&ctx->rx_mu);

    pthread_mutex_lock(&ctx->rx_mu);
    topoll = ctx->recv_offset - ctx->rx_offset;
    pthread_mutex_unlock(&ctx->rx_mu);
    while (topoll == 0) {
        if (connRdmaHandleCq(ctx) != VALKEY_OK) {
            return VALKEY_ERR;
        }
        pthread_mutex_lock(&ctx->rx_mu);
        topoll = ctx->recv_offset - ctx->rx_offset;
        pthread_mutex_unlock(&ctx->rx_mu);
        if (topoll != 0) {
            break;
        }
        if (valkeyRdmaPollCqCm(ctx, end) != VALKEY_OK) {
            return VALKEY_ERR;
        }
    }
    goto pollcq;
}

ssize_t rdmaWrite(RdmaContext *ctx, const char *obuf, size_t data_len, long timeout_msec) {
    struct rdma_cm_id *cm_id = ctx->cm_id;
    long end;
    uint32_t towrite, topoll, wrote = 0;
    size_t ret;
    end = vk_msec_now() + timeout_msec;

pollcq:
    pthread_mutex_lock(&ctx->tx_mu);
    assert(ctx->tx_offset <= ctx->tx_length);
    if (ctx->tx_offset == ctx->tx_length) {
        /* wait a new TX buffer */
        pthread_mutex_unlock(&ctx->tx_mu);
        goto waitcq;
    }

    towrite = valkeyMin(ctx->tx_length - ctx->tx_offset, data_len - wrote);
    ret = connRdmaSend(ctx, cm_id, obuf + wrote, towrite);
    pthread_mutex_unlock(&ctx->tx_mu);
    if (ret == (size_t)VALKEY_ERR) {
        return VALKEY_ERR;
    }

    wrote += ret;
    if (wrote == data_len) {
        return data_len;
    }

waitcq:
    pthread_mutex_lock(&ctx->tx_mu);
    topoll = ctx->tx_offset - ctx->tx_length;
    pthread_mutex_unlock(&ctx->tx_mu);
    while (topoll == 0) {
        if (connRdmaHandleCq(ctx) != VALKEY_OK) {
            return VALKEY_ERR;
        }
        pthread_mutex_lock(&ctx->tx_mu);
        topoll = ctx->tx_offset - ctx->tx_length;
        pthread_mutex_unlock(&ctx->tx_mu);
        if (topoll != 0) {
            break;
        }
        if (valkeyRdmaPollCqCm(ctx, end) != VALKEY_OK) {
            return VALKEY_ERR;
        }
    }
    goto pollcq;
}

void rdmaDisconnect(RdmaContext *ctx) {
    struct rdma_cm_id *cm_id;
    cm_id = ctx->cm_id;
    rdma_disconnect(cm_id);
}

void rdmaClose(RdmaContext *ctx) {
    struct rdma_cm_id *cm_id;

    cm_id = ctx->cm_id;
    connRdmaHandleCq(ctx);
    ibv_destroy_cq(ctx->cq);
    rdmaDestroyIoBuf(ctx);
    ibv_destroy_qp(cm_id->qp);
    ibv_destroy_comp_channel(ctx->comp_channel);
    ibv_dealloc_pd(ctx->pd);
    rdma_destroy_id(cm_id);
    rdma_destroy_event_channel(ctx->cm_channel);

    pthread_mutex_destroy(&ctx->cq_mu);
    pthread_mutex_destroy(&ctx->rx_mu);
    pthread_mutex_destroy(&ctx->tx_mu);
    pthread_mutex_destroy(&ctx->err_mu);
}