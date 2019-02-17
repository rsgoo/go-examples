package main

import (
	"go_dev/day01/package_example/cacl"
	"fmt"
)

func main() {
	a, b := 11, 22
	addVal := cacl.Add(a, b)
	subVal := cacl.Sub(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, addVal)
	fmt.Printf("%v - %v = %v\n", a, b, subVal)
}
