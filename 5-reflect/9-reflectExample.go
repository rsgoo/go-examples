package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s *Student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s *Student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	fmt.Println(v.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("methodName: %s\n", t.Method(i).Name)
		fmt.Printf("method: %s\n", t.Method(i).Type)
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu1 := &Student{
		Name:  "汤姆",
		Score: "90",
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	printMethod(stu1)

}
