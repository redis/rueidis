package rueidis

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"testing/synctest"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

func TestPipeFullRingContext(t *testing.T) {
	for _, multi := range []bool{false, true} {
		name := "Do"
		if multi {
			name = "DoMulti"
		}
		t.Run(name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				p, mock, _, closeConn := setup(t, ClientOption{RingScaleEachConn: 1})
				defer func() {
					closeConn()
					p.Close()
				}()
				p.background()

				// 服务端收下前两条命令但不回复，保持队列饱和。
				pending := make([]chan RedisResult, 2)
				for i := range pending {
					pending[i] = make(chan RedisResult, 1)
					go func() {
						pending[i] <- p.Do(context.Background(), cmds.NewCompleted([]string{"GET", strconv.Itoa(i)}))
					}()
					mock.Expect("GET", strconv.Itoa(i))
				}

				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
				defer cancel()
				done := make(chan []error, 1)
				go func() {
					if multi {
						responses := p.DoMulti(ctx,
							cmds.NewCompleted([]string{"GET", "canceled-1"}),
							cmds.NewCompleted([]string{"GET", "canceled-2"}),
						)
						done <- []error{responses.s[0].Error(), responses.s[1].Error()}
					} else {
						done <- []error{p.Do(ctx, cmds.NewCompleted([]string{"GET", "canceled"})).Error()}
					}
				}()
				synctest.Wait()
				time.Sleep(10 * time.Millisecond)
				synctest.Wait()
				select {
				case errs := <-done:
					for _, err := range errs {
						if !errors.Is(err, context.DeadlineExceeded) {
							t.Errorf("request returned %v, want context.DeadlineExceeded", err)
						}
					}
				default:
					// 失败时先关闭模拟连接，确保旧实现的阻塞调用也能退出。
					closeConn()
					synctest.Wait()
					<-done
					t.Fatal("request did not return when its deadline expired while the ring was full")
				}

				// 恢复回复，验证原请求的响应配对，以及后续请求仍能使用同一连接。
				for i := range pending {
					mock.Expect().ReplyString(strconv.Itoa(i))
					got, err := (<-pending[i]).ToString()
					if err != nil || got != strconv.Itoa(i) {
						t.Fatalf("response %d = (%q, %v)", i, got, err)
					}
				}
				next := make(chan RedisResult, 1)
				go func() {
					next <- p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "next"}))
				}()
				// 若取消的命令仍被发送，这里会读到错误的 key 并失败。
				mock.Expect("GET", "next").ReplyString("next-value")
				got, err := (<-next).ToString()
				if err != nil || got != "next-value" {
					t.Fatalf("next response = (%q, %v)", got, err)
				}
				synctest.Wait()
				if waits := p.loadWaits(); waits != 0 {
					t.Errorf("canceled enqueue leaked %d pending requests", waits)
				}
			})
		})
	}
}
