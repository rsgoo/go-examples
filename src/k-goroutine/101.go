package main

import (
	"fmt"
	"runtime"
	"sync"
)

var myWg sync.WaitGroup

func a() {
	defer myWg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}

func b() {
	defer myWg.Done()
	for i := 10; i < 20; i++ {
		fmt.Println("b:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	myWg.Add(2)
	go a()
	go b()
	myWg.Wait()
	fmt.Println("main")
}
