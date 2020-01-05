package main

import (
	"x-zinx/znet"
	"x-zinx/ziface"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter
}

//test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router handle...")

	fmt.Println("receive from client: msgID = ", request.GetMsgID(), " data = ", string(request.GetData()))

	//先读取客户端的数据，在回写 ping...ping...ping
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

//test Handle
func (pr *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("call HelloZinxRouter handle...")

	fmt.Println("receive from client: msgID = ", request.GetMsgID(), " data = ", string(request.GetData()))

	//先读取客户端的数据，在回写 ping...ping...ping
	err := request.GetConnection().SendMsg(201, []byte("hello...zinx...zinx...zinx"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//1：创建一个Server句柄,使用Zinx的Api
	server := znet.NewServer("【zinx v0.7】")

	//给当前的zinx框架添加自定义router
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloZinxRouter{})

	//2：启动Server
	server.Serve()

}
