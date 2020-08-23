package main

import (
	"fmt"
	"time"
)

func main() {
	//timer := time.NewTimer(5 * time.Second)
	//fmt.Println(time.Now())
	//data := <-timer.C
	//fmt.Printf("%T\n", timer.C)
	//fmt.Printf("%T\n", data)
	//fmt.Println(data)

	timer := time.After(5 * time.Second)
	fmt.Println(time.Now())
	data := <-timer
	fmt.Printf("%T\n", data)
	fmt.Println(data)
}
