package main

import (
	"net"
	"fmt"
	"os"
)

func HandleUdpServerError(err error, when string) {
	if err != nil {
		fmt.Println(when, " : ", err)
		os.Exit(1)
	}
}

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	HandleUdpServerError(err, "server net.ResolveUDPAddr")

	udpConn, err := net.ListenUDP("udp", udpAddr)
	HandleUdpServerError(err, "udp net.ListenUDP")

	defer udpConn.Close()

	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, err := udpConn.ReadFromUDP(buffer)
		HandleUdpServerError(err, "server udpConn.ReadFromUDP")

		clientMsg := string(buffer[:n])
		fmt.Printf("读到来自%v的内容: %s\n", remoteAddr, clientMsg)
		if clientMsg == "bye" {
			udpConn.WriteToUDP([]byte("Time to :"+"Good Bye"), remoteAddr)
		} else {
			udpConn.WriteToUDP([]byte("朕已阅:"+clientMsg), remoteAddr)
		}
	}
	fmt.Println("servr disconnet")

}
