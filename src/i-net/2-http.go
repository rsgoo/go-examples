package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "www.liwenzhou.com")

	if err != nil {
		fmt.Println("访问失败:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn,"get")

}
