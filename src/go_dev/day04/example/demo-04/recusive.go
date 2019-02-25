package demo_04

import (
	"fmt"
	"time"
)

func Recursive(n int) {
	if n > 10 {
		return
	}
	fmt.Println("hello - ", n)
	time.Sleep(time.Nanosecond)

	Recursive(n + 1)
}
