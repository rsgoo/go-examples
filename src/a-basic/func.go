package main

import "fmt"

func adder() func(int) int {
	var x int
	fmt.Println("x = ", x)
	return func(y int) int {
		x += y
		fmt.Println("x == ", x)
		return x
	}
}

func a(language string) (func(name string) string) {
	return func(name string) string {
		return language + " hello,world " + name
	}
}

func main() {
	x := a("php")
	fmt.Println(x("inscode"))
	fmt.Println(x("lisp"))
}
