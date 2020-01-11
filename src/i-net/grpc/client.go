package main

import (
	"google.golang.org/grpc"
	"fmt"
	"i-net/grpc/message"
	"time"
	"golang.org/x/net/context"
)

func main() {

	//连接gPRC服务
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())

	if err != nil {
		fmt.Println("grpc dial err:", err)
		return
	}

	defer conn.Close()

	//将连接绑定到一个服务
	orderServiceClient := message.NewOrderServiceClient(conn)

	//定义请求参数
	orderRequest := message.OrderRequest{
		OrderId:   "20200111002",
		Timestamp: time.Now().Unix(),
	}

	//调用服务的业务方法
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), &orderRequest)

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(orderInfo)
	fmt.Println(orderInfo.GetOrderId())
	fmt.Println(orderInfo.GetOrderName())

}
