package main

import "fmt"

type Car interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

func (bmw *BMW) GetName() string {
	return bmw.Name
}

func (bmw *BMW) Run() {
	fmt.Println(bmw.Name + "can running")
}

func (bmw *BMW) DiDi() {
	fmt.Println(bmw.Name + "can DiDi")
}

func main() {
	var bmw Car

	bmw = &BMW{
		Name: "BMW M6",
	}
	bmwName := bmw.GetName()
	fmt.Println(bmwName)
	bmw.Run()
	bmw.DiDi()
}
