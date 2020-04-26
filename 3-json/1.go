package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Rmb    float64
	Gender string
	Hobby  []string
}

func main() {
	var Tom = Person{
		Name:   "Tom",
		Age:    2,
		Rmb:    268000,
		Gender: "Male",
		Hobby:  []string{"eat fish", "play with tom"},
	}

	bytes, err := json.Marshal(Tom)
	if err != nil {
		fmt.Println("err is ", nil)
	}
	fmt.Println(string(bytes))
}
