package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接服务端失败,err: ", err)
		return
	}

	//向server端发送信息
	buffer := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	for {
		lineString, _ := reader.ReadString('\n')
		lineString = strings.Replace(lineString, "\n", "", -1)

		_, err := conn.Write([]byte(lineString))

		if err != nil {
			fmt.Println("发送给服务端失败:", err)
			return
		}

		m, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取务端失败:", err)
			return
		}

		serverMsg := string(buffer[:m])
		fmt.Println("msg return from server: ", serverMsg)
		if serverMsg == "Chat finish, Good Bye" {
			break;
		}

	}
}
