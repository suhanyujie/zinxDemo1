package main

import (
	"fmt"
	"zinx_study1/znet"
)

func main() {
	fmt.Println("123")
	znet.NewServer("FirstZinx").Serve()
}
