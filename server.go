package main

import (
	"fmt"
	"log"
	"zinx_study1/ziface"
	"zinx_study1/znet"
)

// 框架使用者，定义路由
type UserRouter struct {
	znet.BaseRouter
}

// before business handler
func (_this *UserRouter) PreHandle(request ziface.IRequest) {
	msg := "this is PreHandle\n"
	log.Println(msg)
	request.GetConnection().GetTcpConnection().Write([]byte(msg))
}

// doing business handler
func (_this *UserRouter) DoingHandle(request ziface.IRequest) {
	msg := "this is DoingHandle\n"
	log.Println(msg)
	request.GetConnection().GetTcpConnection().Write([]byte(msg))
}

// after business handler
func (_this *UserRouter) AfterHandle(request ziface.IRequest) {
	msg := "this is AfterHandle\n"
	log.Println(msg)
	request.GetConnection().GetTcpConnection().Write([]byte(msg))
}

func main() {
	fmt.Println("server start...")
	server := znet.NewServer("FirstZinx")
	server.AddRoute(&UserRouter{})
	server.Serve()
}
