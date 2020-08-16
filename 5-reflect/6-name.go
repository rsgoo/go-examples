package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v  kind: %v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32
	var b myInt
	var c rune
	//指针变量的type 为空，kind 为 ptr
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type Person struct {
		name string
		age  int
	}

	type book struct {
		title string
	}

	var d = Person{
		name: "清河",
		age:  10,
	}

	var e = book{title: "《学Golang》"}
	//type: Person  kind: struct
	reflectType(d)
	//type: book  kind: struct
	reflectType(e)
}
