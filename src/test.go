package main

import "fmt"

type Base struct {
}

func (b *Base) Hello() {
	fmt.Println("hello base")
}

type Middle struct {
	Base
}

func (n *Middle) Hello() {
	fmt.Println("hello middle")
}

func main() {
	var s = &Middle{
		Base{},
	}
	s.Hello()
}
