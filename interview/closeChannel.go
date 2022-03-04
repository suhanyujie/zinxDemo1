package main

import (
	"fmt"
	"time"
)

// use a closed channel, what will happened? https://www.jianshu.com/p/e89dfebe2c4a
// > 会发生 panic
func main() {
	whatHappened()
}

func whatHappened() {
	var ch chan string
	go putData(ch)
	time.Sleep(3 * time.Second)
	close(ch)
	select {
	case str := <-ch:
		fmt.Println(str)
	}
	select {}
}

func putData(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "test"
		time.Sleep(1 * time.Second)
	}
}
