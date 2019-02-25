package demo_04

import "fmt"

func SliceCode() {
	var a []int
	a = append(a, 1, 11, 111)
	fmt.Println(a)
	a = append(a, a...)
	fmt.Println("数组 append 一个数组 使用后...", a)
}
