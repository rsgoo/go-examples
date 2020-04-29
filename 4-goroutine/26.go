package main

import (
	"fmt"
	"time"
)

func Hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go Hello()

	for i := 0; i < 10; i++ {
		fmt.Println("main", i)
		time.Sleep(time.Second)
	}
}
