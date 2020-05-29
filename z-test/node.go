package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func main() {
	ce := make(map[int]student)
	ce[1] = student{12, "xiaolizi", 22}
	ce[2] = student{13, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}
