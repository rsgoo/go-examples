package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func(ix int) {
			wg.Add(1)
			fmt.Println("aaa", ix)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("over")
}
