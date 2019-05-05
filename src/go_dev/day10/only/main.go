package main

import "fmt"

func main() {
	//声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 2)
	chan2 <- 1
	chan2 <- 2
	fmt.Println(chan2)

	//声明为只读
	var chan3 <-chan int
	fmt.Println(chan3)

}
