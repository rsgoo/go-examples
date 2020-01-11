package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Printf("输入阶乘数字")
	fmt.Scan(&num)
	fmt.Println(fact(num))
}

func fact(n int) int {
	if n <= 1 {
		return 1
	}

	return n * fact(n-1)
}
