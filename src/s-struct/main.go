package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//var num int64 = 100
	//reflectTest(num)
	var student = Student{
		Name: "Tom-cat",
		Age:  2,
	}
	reflectTest(student)

}

//通过反射获取到传入的类型与类别
func reflectTest(b interface{}) {
	refType := reflect.TypeOf(b)

	fmt.Println("refType:", refType)

	refValue := reflect.ValueOf(b)

	refInterface := refValue.Interface()
	fmt.Printf("refInterface = %v refInterface = %T\n", refInterface)

}
