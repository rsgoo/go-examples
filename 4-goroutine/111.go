package main

import "fmt"

func main() {
	ch1 := make(chan int)

	go func() {
		ch1 <- 100
		ch1 <- 200
		ch1 <- 300
		close(ch1)
		ch1 <- 400
	}()

	data, ok := <-ch1
	fmt.Println(data, ok)

	data, ok = <-ch1
	fmt.Println(data, ok)

	data, ok = <-ch1
	fmt.Println(data, ok)

	data, ok = <-ch1
	fmt.Println(data, ok)
}
