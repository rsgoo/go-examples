package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name  string
	Age   int
	Score float64
}

func (obj Animal) Greeting() {
	fmt.Println("Hello, Jerry")
}

func (obj Animal) Play() {
	fmt.Println("Play with Tom")
}

func main() {
	var HTMLstr = ""
	fmt.Println(HTMLstr)
	tom := Animal{
		Name:  "tom",
		Age:   2,
		Score: 88,
	}
	fmt.Println(reflect.TypeOf(tom))
	fmt.Println(reflect.TypeOf(tom).Kind())
	fmt.Println(reflect.TypeOf(tom).NumField())
	fmt.Println(reflect.TypeOf(tom).NumMethod())
	for i := 0; i < reflect.TypeOf(tom).NumField(); i++ {
		fmt.Println(reflect.TypeOf(tom).Field(i))
		fmt.Println(reflect.TypeOf(tom).Field(i).Name)
		fmt.Println(reflect.TypeOf(tom).Field(i).Index)
		fmt.Println(reflect.TypeOf(tom).Field(i).Type)
		fmt.Println(reflect.TypeOf(tom).Field(i).Anonymous)
		fmt.Println("-----------------------------")
	}

}
