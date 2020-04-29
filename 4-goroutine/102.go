package main

import (
	"fmt"
	"runtime"
	"time"
)

func a1() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b1() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go a1()
	go b1()
	time.Sleep(time.Second)
}
