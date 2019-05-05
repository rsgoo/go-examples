package main

import (
	"fmt"
	"runtime"
)

func add(a, b int, pipe chan int) {
	sum := a + b
	pipe <- sum
}

func main() {
	var pipe chan int
	pipe = make(chan int, 1)
	go add(11, 22, pipe)
	sum := <-pipe
	fmt.Println("sum from pipe is ", sum)
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
}
