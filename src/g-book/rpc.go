package main

import (
	"time"
	"errors"
	"fmt"
)

func main() {
	ch := make(chan string)

	go RPCServer(ch)

	receive, err := PRCClient(ch, "hi")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client receive from server:", receive)
	}
}

func PRCClient(ch chan string, req string) (string, error) {
	ch <- req

	//多路复用，优先匹配有数据返回的通道
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second): //超时1秒匹配
		return "", errors.New("time out")
	}
}

func RPCServer(ch chan string) {
	for {
		//接收客户端请求
		data := <-ch

		fmt.Println("server receive from client:", data)

		//模拟超时
		//time.Sleep(time.Second * 2)

		ch <- "roger that, move out!"
	}
}
