package main

import (
	"net"
	"fmt"
	"os"
)

func HandleServerError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func dealWithConn(conn net.Conn) {
	//准备缓冲区
	buffer := make([]byte, 1024)

	for {
		//读取客户端连接时传递给sever的数据,并写到缓冲区buffer中
		n, err := conn.Read(buffer)
		HandleServerError(err, "conn.Read")
		clientMsg := string(buffer[:n]) + " <<from server>>"
		fmt.Printf("received client: %v msg:%s\n", conn.RemoteAddr(), clientMsg)

		//客户端下线处理
		if clientMsg == "im off" {
			conn.Write([]byte("Good Bye"))
			break
		}

		//将读取到的客户端连接数据拼装处理后返回
		conn.Write([]byte("msg received: " + clientMsg))
	}
	fmt.Println("client disconnected...")
}

func main() {
	//建立监听
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	HandleServerError(err, "net.Listen")

	fmt.Println("listening...")

	for {
		//接收客户端连接请求
		conn, err := listener.Accept()
		HandleServerError(err, "listener.Accept")

		//处理客户端连接
		go dealWithConn(conn)

	}
}
