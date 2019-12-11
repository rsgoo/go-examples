package main

import (
	"net"
	"fmt"
)

func ServerBack(address string, exitChan chan int) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(" net.Listen err:", err.Error())
		exitChan <- 1
	}
	defer listener.Close()

	fmt.Println("listen " + listener.Addr().String())
	fmt.Println("listen " + address)

	for {
		connector, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err.Error())
			continue
		}
		go handleSession(connector, exitChan)
	}
}
