package t1

import (
	"fmt"
	"testing"
	"time"
)

// 调用 a, b 函数，保证打印的数据，和原切片的顺序一致
func TestT1(t *testing.T) {
	s1 := []string{"a", "b", "c", "d", "e"}
	ch1 := make(chan int)
	ch2 := make(chan int)
	endCh := make(chan struct{})
	go func() {
		for i, _ := range s1 {
			ch1 <- i
		}
		endCh <- struct{}{}
	}()
	go b(s1, ch2)
	go a(s1, ch1, ch2)

	// 阻塞主协程，使工作协程执行完
	time.Sleep(2 * time.Second)
}

// 只可打印奇数下标的值
func a(p []string, ch1 chan int, ch2 chan int) {
	for {
		select {
		case index := <-ch1:
			if index%2 != 0 {
				fmt.Println(p[index])
			} else {
				ch2 <- index
			}
		}
	}
}

// 只可打印偶数下标的值
func b(p []string, ch1 chan int) {
	for {
		select {
		case index := <-ch1:
			if index%2 == 0 {
				fmt.Println(p[index])
			}
		}
	}
}
