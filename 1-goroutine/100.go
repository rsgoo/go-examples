package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hi(i int) {
	defer wg.Done()
	fmt.Println("hi! ", i)
}

func main() {
	defer fmt.Println("defer func")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Hi(i)
	}
	wg.Wait()
	fmt.Println("main over")

}
