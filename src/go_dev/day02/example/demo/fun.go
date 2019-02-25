package demo

import "fmt"

type Add_func func(int, int) int

func Add(a, b int) int {
	return a + b
}

func compute(option Add_func, a, b int) int {
	return option(a, b)
}

func GetSum() {
	c := Add
	fmt.Println(c)
	sum := compute(c, 110, 220)
	fmt.Printf("%d %v %f = %d", 100, c, 220, sum)
}
