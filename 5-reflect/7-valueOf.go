package main

import (
	"fmt"
	"reflect"
)

func reflectValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(2000)
	}
}

func main() {
	var a int64 = 100
	reflectValue2(&a)
	fmt.Println(a)
}
