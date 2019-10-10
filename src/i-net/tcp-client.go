package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
)

func HandleClientError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	//拨号连接到服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	HandleClientError(err, "net.Dial")

	//建立读取器
	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		//读取一行
		lineBytes, _, err := reader.ReadLine()
		HandleClientError(err, "reader.ReadLine")

		//将client读取到的一行数据写给连接的server
		conn.Write(lineBytes)
		fmt.Println("client input content is :", string(lineBytes))

		//读取连接数据到缓冲区
		n, err := conn.Read(buffer)
		HandleClientError(err, "conn.Read")

		serverMsg := string(buffer[:n])
		fmt.Println("server: ", serverMsg)

	}

	fmt.Println("game finish")
}
