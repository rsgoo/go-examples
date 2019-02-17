package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main func start")
	for i := 0; i < 100; i++ {
		go test_goroutine(i)
	}

	time.Sleep(time.Second)

	fmt.Println("main func exit")
}
