package main

import (
	"net"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen : 2000 失败。err：", err)
		return
	}

	for {
		//如果没有客户端的请求会处于阻塞状态
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("listener failed！err :", err)
			return
		}

		//使用协程处理客户端请求的连接
		go process(conn)
	}
}

func process(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取client数据错误：", err)
		}
		clientMsg := string(buffer[:n])
		fmt.Println("请求来自：", conn.RemoteAddr())
		fmt.Println("读取client的数据是：", clientMsg)
		conn.RemoteAddr().String()
		conn.Write([]byte("这是来自服务端的响应：" + clientMsg))
	}
}
