package main

import "fmt"

func main() {
	n := 9
	fmt.Println(sumNums(n))
}

func sumNums(n int) int {
	if n <= 1 {
		return 1
	}
	return n + sumNums(n-1)
}
