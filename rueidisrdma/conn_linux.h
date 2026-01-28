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

#ifndef VALKEY_GO_RDMA_H
#define VALKEY_GO_RDMA_H

#define VALKEY_RDMA_MAX_WQE 1024
#define VALKEY_RDMA_DEFAULT_RX_LEN (1024 * 1024)
#define VALKEY_RDMA_INVALID_OPCODE 0xffff

#define VALKEY_ERR -1
#define VALKEY_OK 0

/* When an error occurs, the err flag in a context is set to hold the type of
 * error that occurred. VALKEY_ERR_IO means there was an I/O error and you
 * should use the "errno" variable to find out what is wrong.
 * For other values, the "errstr" field will hold a description. */
#define VALKEY_ERR_IO 1       /* Error in read or write */
#define VALKEY_ERR_EOF 3      /* End of file */
#define VALKEY_ERR_PROTOCOL 4 /* Protocol error */
#define VALKEY_ERR_OOM 5      /* Out of memory */
#define VALKEY_ERR_TIMEOUT 6  /* Timed out */
#define VALKEY_ERR_OTHER 2    /* Everything else... */

/* Connection may be disconnected before being free'd. The second bit
 * in the flags field is set when the context is connected. */
#define VALKEY_CONNECTED 0x2

#include <pthread.h>
#include <rdma/rdma_cma.h>

typedef struct valkeyRdmaFeature {
    /* defined as following Opcodes */
    uint16_t opcode;
    /* select features */
    uint16_t select;
    uint8_t rsvd[20];
    /* feature bits */
    uint64_t features;
} valkeyRdmaFeature;

typedef struct valkeyRdmaKeepalive {
    /* defined as following Opcodes */
    uint16_t opcode;
    uint8_t rsvd[30];
} valkeyRdmaKeepalive;

typedef struct valkeyRdmaMemory {
    /* defined as following Opcodes */
    uint16_t opcode;
    uint8_t rsvd[14];
    /* address of a transfer buffer which is used to receive remote streaming data,
     * aka 'RX buffer address'. The remote side should use this as 'TX buffer address' */
    uint64_t addr;
    /* length of the 'RX buffer' */
    uint32_t length;
    /* the RDMA remote key of 'RX buffer' */
    uint32_t key;
} valkeyRdmaMemory;

typedef union valkeyRdmaCmd {
    valkeyRdmaFeature feature;
    valkeyRdmaKeepalive keepalive;
    valkeyRdmaMemory memory;
} valkeyRdmaCmd;

typedef enum valkeyRdmaOpcode {
    GetServerFeature = 0,
    SetClientFeature = 1,
    Keepalive = 2,
    RegisterXferMemory = 3,
} valkeyRdmaOpcode;

typedef struct RdmaContext {
    struct rdma_cm_id *cm_id;
    struct rdma_event_channel *cm_channel;
    struct ibv_comp_channel *comp_channel;
    struct ibv_cq *cq;
    struct ibv_pd *pd;

    /* TX */
    char *tx_addr;
    uint32_t tx_length;
    uint32_t tx_offset;
    uint32_t tx_key;
    char *send_buf;
    uint32_t send_length;
    uint32_t send_ops;
    struct ibv_mr *send_mr;

    /* RX */
    uint32_t rx_offset;
    char *recv_buf;
    unsigned int recv_length;
    unsigned int recv_offset;
    struct ibv_mr *recv_mr;

    /* CMD 0 ~ VALKEY_RDMA_MAX_WQE for recv buffer
     * VALKEY_RDMA_MAX_WQE ~ 2 * VALKEY_RDMA_MAX_WQE -1 for send buffer */
    valkeyRdmaCmd *cmd_buf;
    struct ibv_mr *cmd_mr;

    int flags;
    int err;          /* Error flags, 0 when there is no error */
    char errstr[64];  /* String representation of error when applicable */

    pthread_mutex_t cq_mu;
    pthread_mutex_t rx_mu;
    pthread_mutex_t tx_mu;
    pthread_mutex_t err_mu;

} RdmaContext;

#endif /* VALKEY_RDMA_H */
