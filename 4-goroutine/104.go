package main

import "fmt"

func receive(ch chan bool) {
	result := <-ch

	fmt.Println(result)
}

func main() {
	ch := make(chan bool)
	go receive(ch)
	ch <- true
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
	fmt.Println("main over")
}
