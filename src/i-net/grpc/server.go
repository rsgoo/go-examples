package main

import (
	"golang.org/x/net/context"
	"i-net/grpc/message"
	"time"
	"errors"
	"google.golang.org/grpc"
	"net"
	"fmt"
)

type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
	orderMap := map[string]message.OrderInfo{
		"20200111001": {
			OrderId:     "20200111001",
			OrderName:   "手机",
			OrderStatus: "已付款",
		},
		"20200111002": {
			OrderId:     "20200111002",
			OrderName:   "电脑",
			OrderStatus: "已付款",
		},
		"20200111003": {
			OrderId:     "20200111003",
			OrderName:   "平板",
			OrderStatus: "未付款",
		},
	}

	var response message.OrderInfo

	current := time.Now().Unix()
	if request.Timestamp > current {
		response = message.OrderInfo{
			OrderId:     "0",
			OrderName:   "",
			OrderStatus: "订单信息移除",
		}
	} else {
		response = orderMap[request.OrderId]
		if response.OrderId != "" {
			response = orderMap[request.OrderId]
		} else {
			return nil, errors.New("server error")
		}
	}

	return &response, nil
}

func main() {
	//声明一个gRPC服务
	server := grpc.NewServer()

	//注册一个服务业务
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	//监听
	listen, err := net.Listen("tcp", ":8088")

	if err != nil {
		fmt.Println("err info: ", err)
		return
	}

	//启动一个服务
	server.Serve(listen)

}
