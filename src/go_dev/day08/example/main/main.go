package main

import "fmt"

type Shaper interface {
	Area() float32
	Perimeter() float32
}
type Square struct {
	side float32
}

func (this *Square) Area() float32 {
	return this.side * this.side
}

func (this *Square) Perimeter() float32 {
	return this.side * 4
}

type Stringer interface {
	String() string
}

func main() {

	var v Stringer

	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	} else {
		fmt.Println("a")
	}
}
