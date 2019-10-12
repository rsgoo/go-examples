package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"strings"
)

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}

}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	ClientHandleError(err, "client net.Dial")
	buffer := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)

	for {
		lineString, _ := reader.ReadString('\n')
		lineString = strings.Replace(lineString, "\n", "", -1)

		if len(lineString) == 0 {
			fmt.Println("聊天内容不能为空哦")
			continue
		}

		n, err := conn.Write([]byte(lineString))
		ClientHandleError(err, "client conn.Write")

		fmt.Printf("%d 发送的消息是: %v\n", n, lineString)

		m, err := conn.Read(buffer)
		ClientHandleError(err, "client conn.Read")

		serverMsg := string(buffer[:m])
		fmt.Println("msg return from server: ", serverMsg)
		if serverMsg == "Chat finish, Good Bye" {
			break;
		}
	}

	fmt.Println("game finish")
}
