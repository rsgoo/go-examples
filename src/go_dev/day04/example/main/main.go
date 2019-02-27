package main

import (
	"go_dev/day04/example/demo-04"
	"fmt"
)

func main() {
	fmt.Println("test func new")
	demo_04.NewCode()

	fmt.Println("test func slice")
	demo_04.SliceCode()

	fmt.Println("test func recursive")
	demo_04.Recursive(0)

	fmt.Println("test func closure")
	demo_04.Closure()

	fmt.Println("--------------------------")

	s1 := make([]int, 10)
	s1 = []int{1, 2, 3, 4, 5, 6, 6, 7}
	fmt.Println(s1)

	s2 := s1[:len(s1)]
	fmt.Println(s2)

	demo_04.TestSlice()
}
