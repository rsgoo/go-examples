package main

import "fmt"

func main() {
	var myIntChan chan int
	fmt.Println(myIntChan) //nil

	//给管道开辟一块内存地址
	myIntChan = make(chan int, 3)
	fmt.Println(myIntChan)

	myIntChan <- 1
	myIntChan <- 2
	myIntChan <- 3

	for value := range myIntChan {
		fmt.Println(value)
	}
}
