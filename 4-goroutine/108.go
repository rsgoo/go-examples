package main

import "fmt"

func main() {
	ch := make(chan int)
	go goRun(ch)
	ch <- 100

	fmt.Println("hello,golang")
}

func goRun(ch chan int) {
	fmt.Println(<-ch)
}
