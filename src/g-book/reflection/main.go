package main

import (
	"reflect"
	"fmt"
)

func main() {
	var num int64 = 1024
	reflectTest(num)
}

//通过反射获取到传入的类型与类别
func reflectTest(b interface{}) {
	refVal := reflect.ValueOf(b)
	fmt.Printf("%v,%T\n", refVal, refVal)
	fmt.Printf("%v,%T\n", refVal.Interface(), refVal.Interface().(int))
}
