package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

func handleSessionBackup(connector net.Conn, exitChan chan int) {
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
