package main

import (
	"net"
	"fmt"
	"os"
)

func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}

}

func ChatWith(conn net.Conn) {

	buffer := make([]byte, 1024)
	for {
		//将读取的内容放置在缓冲区中
		n, err := conn.Read(buffer)
		ServerHandleError(err, "server conn.Read")
		clientMsg := string(buffer[:n])

		fmt.Printf("收到 %s 发送的消息:%v <>\n", conn.RemoteAddr(), clientMsg)

		if clientMsg != "bye" {

			if clientMsg == "rust" || clientMsg == "lisp" || clientMsg == "golang" {
				clientMsg = "你想学习" + clientMsg + "吗？"
			}
			//将读取到的客户端连接数据拼装处理后返回
			conn.Write([]byte("消息已阅：" + clientMsg))
		} else {
			conn.Write([]byte("Chat finish, Good Bye"))
			//这里使用break不断开server端
			break
		}
	}
	conn.Close()
	fmt.Printf("客户端 %s 已经断开连接\n", conn.RemoteAddr())

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	ServerHandleError(err, "server net.Listen")
	fmt.Println("server start listening...")
	for {

		conn, err := listener.Accept()

		ServerHandleError(err, "server listener.Accept")

		//接入client连接并进行处理
		go ChatWith(conn)
	}

}
