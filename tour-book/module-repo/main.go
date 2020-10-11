package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	b := &a
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
	b.Name = "apple"
	fmt.Println(a)
	fmt.Println(*b)
}
