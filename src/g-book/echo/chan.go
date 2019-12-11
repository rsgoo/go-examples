package main

import "fmt"

func main() {
	myChan := make(chan int)

	//go func() {
	//	myChan <- 1
	//}()

	fmt.Println(<-myChan)
}
