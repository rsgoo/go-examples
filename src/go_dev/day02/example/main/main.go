package main

import (
	"fmt"
	"time"
	"go_dev/day02/example/demo"
)

func main() {
	fmt.Println(time.Month(1).String())
	a := 10
	b := 20
	demo.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
