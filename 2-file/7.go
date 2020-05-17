package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	src, _ := os.OpenFile("2.png", os.O_RDONLY, 0666)
	dst, _ := os.OpenFile("212.png", os.O_WRONLY|os.O_CREATE, 0666)
	n, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
