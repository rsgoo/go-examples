package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) Set(name string, age int, score float32) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func testStruct(a interface{}) {
	val := reflect.TypeOf(a)
	kind := val.Kind()
	if kind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	method := val.NumMethod()
	fmt.Println(num)
	fmt.Println(method)
}

func main() {
	var a = Student{
		Name:  "linus",
		Age:   56,
		Score: 1024,
	}
	a.Set("inscode",23,23.45)
	testStruct(a)
}
