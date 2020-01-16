package main

import (
	"fmt"
	"time"
)

func main() {
	var arrA = []byte{'a', 'b', 'c', 'd'}
	var arr1 = []int{1, 2, 3, 4}

	go func() {
		for _, value := range arrA {
			fmt.Println(value)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for _, value := range arr1 {
			fmt.Println(value)
			time.Sleep(time.Second * 2)
		}
	}()

}
