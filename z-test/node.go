package main

import "fmt"

func main() {
	var str = "hello,world"
	foo := func() {
		str = "hello,foo"
	}
	fmt.Println(str)
	foo()
	fmt.Println(str)

}
