package main

import (
	"fmt"
	"time"
)

func main() {
	go CountNum1()
	CountNum2()
	fmt.Println("main func over")
}

func CountNum1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("countNum1-->", i)
	}
}

func CountNum2() {
	for i := 101; i <= 110; i++ {
		fmt.Println("countNum2-->", i)
		time.Sleep(time.Microsecond)
	}
}
