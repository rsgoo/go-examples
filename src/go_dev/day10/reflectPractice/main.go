package main

import (
	"fmt"
	"reflect"
)

//定义一个结构体
type Monster struct {
	Name  string  `json:"name"`
	Ages  int     `json:"age"`
	Score float32 `json:"score"`
	Sex   string  `json:"sex"`
	//Lang  string
}

func (self Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (self Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(self)
	fmt.Println("----end----")
}

func (self Monster) Set(name string, age int, score float32, sex string) {
	self.Name = name
	self.Ages = age
	self.Score = score
	self.Sex = sex
}

func TestStruct(b interface{}) {
	refType := reflect.TypeOf(b)
	refValue := reflect.ValueOf(b)
	refKind := refType.Kind()
	fmt.Println("refType:", refType)
	fmt.Println("refValue:", refValue)
	fmt.Println("refKind:", refKind)
	if refKind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	fmt.Println("------------------")

	//获取结构体的字段
	structNumFields := refValue.NumField()
	fmt.Println("NumFields", structNumFields)
	for i := 0; i < structNumFields; i++ {
		fmt.Printf("field %d: 值 = %v\n", i, refValue.Field(i))
		//获取标签Tag的值【注意这里是refType】
		tagVal := refType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Println("tag value is: ", tagVal)
			//获取结构体的Name
			fmt.Printf("Name = %v; Type = %v\n", refType.Field(i).Name, refType.Field(i).Type)
			fmt.Println("*****************")
		}
	}
	numOfMethod := refType.NumMethod()
	fmt.Println("numOfMethod: ", numOfMethod)
	//注意这里调用的是refValue; call()中为调用的参数
	refValue.Method(1).Call(nil)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(50))
	//参数调用
	result := refValue.Method(0).Call(params)
	fmt.Println(result[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Ages:  400,
		Score: 30.8,
		Sex:   "male",
	}
	TestStruct(a)
}
