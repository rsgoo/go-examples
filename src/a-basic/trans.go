package main

import "fmt"

type Flier interface {
	Fly()
}
type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	for name, obj := range animals {
		f, isFlier := obj.(Flier)
		w, isWalk := obj.(Walker)

		fmt.Println(name)
		if isFlier {
			f.Fly()
		}

		if isWalk {
			w.Walk()
		}
	}

	p1 := new(pig)
	var a Walker = p1
	p2 := a.(*pig)
	fmt.Println(p2.Walk)
}
