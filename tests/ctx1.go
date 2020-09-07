package main

import (
	"context"
	"log"
	"sync/atomic"
	"time"
)

func main() {
	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	go func(ctx context.Context) {
		var i int64 = 0
		childCtx, childCancel := context.WithCancel(ctx)
		for {
			select {
			case <-ctx.Done():
				childCancel()
				log.Println("ctx done, and childCancel executed")
				return
			default:
				time.Sleep(time.Millisecond * 300)
				if x := atomic.AddInt64(&i, 1); x < 10 {
					go func(c context.Context) {
						for {
							select {
							case <-c.Done():
								log.Printf("child %d end", i)
								return
							default:
								time.Sleep(1 * time.Second)
								log.Printf("child %d running", i)
							}
						}
					}(childCtx)
				}
			}
		}
	}(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	log.Println("main cancel has executed")
	time.Sleep(time.Second * 3)
}
