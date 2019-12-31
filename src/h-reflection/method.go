package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" age:"11019"`
	Score int    `json:"score" grade:"88"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s Student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s Student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	fmt.Println(v.NumMethod())
	fmt.Println("---------------")
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu1 := Student{
		Name:  "小王子",
		Score: 90,
	}

	printMethod(stu1)
}
