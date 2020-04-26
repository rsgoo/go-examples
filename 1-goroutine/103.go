package main

import "fmt"

func main() {
	var ch1 chan int
	var ch2 chan string
	fmt.Println("ch1:", ch1)
	fmt.Println("ch2:", ch2)
	ch1 = make(chan int,1)
	ch1 <- 11
	fmt.Println(<-ch1)
}
