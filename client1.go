package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("client start ...")
	time.Sleep(1 * time.Second)
	// 连接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:3001")
	if err != nil {
		log.Printf("client start error: %s\n", err)
		return
	}
	for {
		// 调用 write 写数据
		_, err = conn.Write([]byte("hello world"))
		if err != nil {
			log.Printf("client write data error: %s\n", err)
			continue
		}
		// 调用完 wirte 后，可以接着从连接中读取数据
		buf := make([]byte, 512)
		_, err = conn.Read(buf)
		if err != nil {
			log.Printf("client read data error: %s\n", err)
			continue
		}
		log.Printf(string(buf))
		time.Sleep(2 * time.Second)
	}
}
