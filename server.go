package main

import (
	"fmt"
	"zinx_study1/znet"
)

func main() {
	fmt.Println("server start...")
	znet.NewServer("FirstZinx").Serve()
}
