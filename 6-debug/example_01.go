package main

import "fmt"

type Animal struct {
	Name  string
	Color string
}

func main() {
	tom := Animal{
		Name:  "tom",
		Color: "blue",
	}
	fmt.Printf("%v\n", tom)
	fmt.Printf("%+v\n", tom)
	fmt.Printf("%#v\n", tom)
}
