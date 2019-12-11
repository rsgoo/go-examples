package main

import (
	"fmt"
	"runtime"
)

func main() {
	myChan := make(chan string)

	for {
		var inputStr string
		fmt.Scan(&inputStr)

		go consumer(myChan)

		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}

func consumer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("consumer data is: ", data)
	}
}
