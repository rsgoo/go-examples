package main

import (
	"math"
	"net/rpc"
	"fmt"
	"net"
	"net/http"
)

type MathUtil struct {
}

// 对外提供服务的方法
func (mu *MathUtil) CalculateCircleArea(req float32, rsp *float32) error {
	*rsp = math.Pi * req * req
	fmt.Println("req = ", req, " result = ", *rsp)
	return nil
}

func main() {
	//0：初始化指针数据类型
	//mathUtil :=new(MathUtil)
	mathUtil := &MathUtil{}

	//1：对服务对象进行注册
	err := rpc.Register(mathUtil)
	if err != nil {
		fmt.Println("err is: ", err)
		return
	}

	//2：通过函数将mathUtil提供的服务注册到http协议上，方便调用者利用http进行调用
	rpc.HandleHTTP()

	//3：在特定的端口进行监听
	listen, err := net.Listen("tcp", ":8081")

	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}

	//4：启动服务
	http.Serve(listen, nil)
}
