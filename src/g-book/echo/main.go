package main

import (
	"os"
	"net"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	exitChan := make(chan int)
	go server("127.0.0.1:7001", exitChan)
	//通道阻塞，等待接收返回值
	code := <-exitChan
	os.Exit(code)

}

func server(address string, exitChan chan int) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(" net.Listen err:", err.Error())
		exitChan <- 1
	}
	defer listener.Close()

	fmt.Println("listen " + listener.Addr().String())

	for {
		connector, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err.Error())
			continue
		}
		go handleSession(connector, exitChan)
	}
}

func handleSession(connector net.Conn, exitChan chan int) {
	fmt.Println("connect session start...")

	//创建一个缓冲读取器
	reader := bufio.NewReader(connector)

	for {
		str, err := reader.ReadString('\n');
		if err != nil {
			fmt.Println("session err: ", err.Error())
			connector.Close()
			break
		}

		str = strings.TrimSpace(str)

		if !processTelnetCommand(str, exitChan) {
			connector.Close()
			break
		}

		connector.Write([]byte(str + "\r\n"))
	}
}

func processTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("server shutdown")
		exitChan <- 0
		return false
	} else {
		fmt.Println(str)
		return true
	}
}
