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
	// request.GetConnection().GetTcpConnection().Write([]byte(msg))
	request.GetConnection().SendMsg(1, []byte(msg))
}

// doing business handler
func (_this *UserRouter) DoingHandle(request ziface.IRequest) {
	// msg := "this is DoingHandle\n"
	msg := request.GetData()
	log.Printf("reveive data from client: %s.\n", string(msg))
	// 获取消息
	request.GetConnection().SendMsg(1, []byte(msg))
}

// 测试路由2
type ExampleRoute1 struct {
	znet.BaseRouter
}

func (_this *ExampleRoute1) DoingHandle(request ziface.IRequest) {
	msg := request.GetData()
	log.Printf("reveive data from client: %s.\n", string(msg))
	// 获取消息
	request.GetConnection().SendMsg(2, []byte("ExampleRoute1: "+string(msg)))
}

// after business handler
func (_this *UserRouter) AfterHandle(request ziface.IRequest) {
	msg := "this is AfterHandle\n"
	log.Println(msg)
	request.GetConnection().SendMsg(1, []byte(msg))
}

func OnStart(connection ziface.IConnection) {
	log.Printf("server: call conn onStart. conn id is: %d", connection.GetConnId())
}

func OnStop(connection ziface.IConnection) {
	log.Printf("server: call conn onStop. conn id is: %d", connection.GetConnId())
}

func main() {
	fmt.Println("server start...")
	// 实例化服务器
	server := znet.NewServer("FirstZinx")
	server.AddRoute(1, &UserRouter{})
	server.AddRoute(2, &ExampleRoute1{})
	// 注册连接开始、关闭时的 hook
	server.SetOnConnStartFunc(OnStart)
	server.SetOnConnStartFunc(OnStop)
	server.Serve()
}
