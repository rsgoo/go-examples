package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func HandleUdpClientError(err error, when string) {
	if err != nil {
		fmt.Println(when, " : ", err)
		os.Exit(1)
	}
}

func main() {
	udpConn, err := net.Dial("udp", "127.0.0.1:8888")
	HandleUdpClientError(err, "client net.Dial")

	defer udpConn.Close()

	buffer := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	localAddr := fmt.Sprintf("%v", udpConn.LocalAddr())

	for {
		clientStr, err := reader.ReadString('\n')
		HandleUdpClientError(err, "client reader.ReadString")

		clientStr = strings.Replace(clientStr, "\n", "", -1)

		n, err := udpConn.Write([]byte(clientStr))
		HandleUdpClientError(err, "client udpConn.Write")

		fmt.Printf("成功写入%d个字节\n", n)

		m, err := udpConn.Read(buffer)
		HandleUdpClientError(err, "udpConn.Read")

		serverMsg := string(buffer[:m])
		if serverMsg == "Time to :Good Bye" {
			fmt.Println("time to say good bye")
			break
		}

		fmt.Println("读取到来自UDP server 的返回:-->", serverMsg)
	}
	fmt.Println(localAddr+" 的client链接断开")
}
