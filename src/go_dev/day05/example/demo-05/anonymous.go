package demo_05

import (
	"time"
	"fmt"
)

type Car struct {
	name string
	age  int
}

type Train struct {
	Car
	int
	start time.Time
}

func GetInfo()  {
	var t Train
	t.Car.name = "train"
	t.Car.age = 11
	t.int = 11019

	fmt.Println(t)
}