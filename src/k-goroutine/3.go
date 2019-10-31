package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 1000; i++ {
		go DoTask(i)
	}
	time.Sleep(time.Second*5)
	fmt.Println("main over")
}

func DoTask(no int) {
	for i := 1; i <= no; i++ {
		fmt.Println("DoTask ", i)
		//time.Sleep(time.Millisecond * 500)
	}
}
