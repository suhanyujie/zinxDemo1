package main

import (
	"fmt"
	"time"
)

// use a closed channel, what will happened? https://www.jianshu.com/p/e89dfebe2c4a
// > 如果只是从一个 closed 的 channel 中读取，则不会报错，只不过读出的值是空值。
// 但如果向一个 closed 的 channel 中 send，则会发生 panic
func main() {
	whatHappened()
}

func whatHappened() {
	var ch = make(chan string, 1)
	go putData(ch)
	time.Sleep(3 * time.Second)
	close(ch)
	select {
	case str := <-ch:
		fmt.Println(str)
	}
	val1 := <-ch
	fmt.Printf("get value: %+v, type: %T\n", val1, val1)
	// select {}
}

func putData(ch chan<- string) {
	ch <- "test"
	time.Sleep(1 * time.Second)
}
