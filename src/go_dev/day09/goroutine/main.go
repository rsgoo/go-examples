package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hello, goroutine: " + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()
	for j := 0; j < 10; j++ {
		fmt.Println("hello main: " + strconv.Itoa(j))
		time.Sleep(time.Second)
	}
}
