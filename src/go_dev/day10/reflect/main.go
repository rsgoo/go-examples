package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
	Lang string
}

func reflectTest01(b interface{}) {
	refType := reflect.TypeOf(b)
	fmt.Println("type is ", refType)
	refValue := reflect.ValueOf(b)
	fmt.Println("value is ", refValue)

	//类型转换
	fmt.Println(refValue.Int() + 1)

	iv := refValue.Interface()

	//使用类型断言
	fmt.Println(iv.(int) + 10) //11029
}

func reflectTest02(b interface{}) {
	refType := reflect.TypeOf(b)
	refValue := reflect.ValueOf(b)
	iv := refValue.Interface()
	fmt.Println(refType)
	fmt.Printf("iv = %v; iv type = %T\n", iv, iv)

	//带检测的类型断言
	stu, ok := iv.(Student)
	if ok {
		fmt.Println("student name:", stu.Name)
		fmt.Println("student age:", stu.Age)
		fmt.Println("student lang:", stu.Lang)
	}
	fmt.Println("value kind is:", refValue.Kind())
	fmt.Println("type kind is:", refType.Kind())
}

//通过反射，修改num int 的值
func reflectTest03(b interface{}) {
	refValue := reflect.ValueOf(b)
	fmt.Println(refValue.Kind())
	//通过反射直接改变值
	refValue.Elem().SetInt(11019)
}

func main() {
	var num int = 19
	//var person = Student{
	//	Name: "aaron",
	//	Age:  26,
	//	Lang: "python",
	//}
	//reflectTest01(person)
	//reflectTest02(person)
	reflectTest03(&num)
	fmt.Println(num)
}
