package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	timer := time.NewTimer(time.Second * 3)

	var i int
	for {
		select {
		case <-timer.C:
			fmt.Println("执行结束")
			goto StopHere
			//return
		case <-ticker.C:
			i++
			fmt.Println("ticker ", i)

		}
	}
StopHere:
	fmt.Println("main over")
}
