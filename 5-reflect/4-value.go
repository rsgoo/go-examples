package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *string

	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
	fmt.Println("var a *int:", reflect.ValueOf(a).IsZero())
	fmt.Println("var a *int:", reflect.ValueOf(a).IsValid())

	var s = struct{}{}
	fmt.Println(reflect.ValueOf(s).MethodByName("").IsValid())

	var m = map[int]int{}
	fmt.Println(reflect.ValueOf(m).MapIndex(reflect.ValueOf(2)).IsValid())
}
