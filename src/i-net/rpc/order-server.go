package main

import (
	"i-net/rpc/message"
	"time"
	"errors"
	"net/rpc"
	"net"
	"fmt"
	"net/http"
)

type OrderService struct {
}

func (os *OrderService) GetOrderInfo(req message.OrderRequest, rsp *message.OrderInfo) error {
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
	current := time.Now().Unix()
	if req.Timestamp > current {
		*rsp = message.OrderInfo{
			OrderId:     "0",
			OrderName:   "",
			OrderStatus: "订单信息移除",
		}
		return nil
	}

	result := orderMap[req.OrderId]
	if result.OrderId != "" {
		*rsp = orderMap[req.OrderId]
	} else {
		return errors.New("server error")
	}

	return nil
}

func main() {
	orderService := &OrderService{}
	rpc.Register(orderService)

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}

	http.Serve(listen, nil)
}
