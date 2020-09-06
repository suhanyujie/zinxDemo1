package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	var doneChan chan bool
	defer1()
	for {
		select {
		case val := <-doneChan:
			log.Printf("%s\n", val)
		default:
			log.Printf("go on waiting\n")
			time.Sleep(1 * time.Second)
		}
	}
}

// go defer 先进后出，后进先出
// 它们会以逆序执行（类似栈，即后进先出）
func defer1() {
	go func() {
		for i := 0; i < 4; i++ {
			defer fmt.Println(i)
		}
	}()
	fmt.Println("end")
	// 为何这里需要 sleep，否则 上面的 defer 看不到执行效果
	// 答：主协程可能很快就退出子协程还未来得及执行
	// time.Sleep(1 * time.Second)
}
