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
	server := znet.NewServer("【zinx v0.9】")

	//注册连接的hook函数
	server.SetOnConnStart(DoConnectionBegin)
	server.SetOnConnStop(DoConnectionAfter)

	//给当前的zinx框架添加自定义router
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloZinxRouter{})

	//2：启动Server
	server.Serve()

}

//创建连接之后的构造函数
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("===>DoConnectionBegin is called<====")
	err := conn.SendMsg(202, []byte("DoConnectionBegin"))
	if err != nil {
		fmt.Println(err)
	}

	//给当前的连接设置一些属性
	fmt.Println("Set conn Name Hooah...")
	conn.SetProperty("name", "雨醉风尘")
	conn.SetProperty("github", "inscode.github.io")
	conn.SetProperty("language", "golang")
}

//断开连接之前的构造函数
func DoConnectionAfter(conn ziface.IConnection) {
	fmt.Println("===>DoConnectionAfter is called<====")
	fmt.Println("connId = ", conn.GetConnID(), " is lost")

	if name, err := conn.GetProperty("name"); err == nil {
		fmt.Println("【【Name = ", name)
	}

	if github, err := conn.GetProperty("github"); err == nil {
		fmt.Println("【【github = ", github)
	}
	if language, err := conn.GetProperty("language"); err == nil {
		fmt.Println("【【language = ", language)
	}

}
