package main

import (
	"net/rpc"
	"fmt"
	"time"
	"i-net/rpc/message"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("dialhttp err:", err)
		return
	}

	timestamp := time.Now().Unix()
	request := message.OrderRequest{
		OrderId:   "20200111001",
		Timestamp: timestamp,
	}

	var response message.OrderInfo

	err = client.Call("OrderService.GetOrderInfo", request, &response)
	if err != nil {
		fmt.Println("call err:", err)
		return
	}
	fmt.Println(response.GetOrderId())
	fmt.Println(response)
}
