package main

import (
	"fmt"
	"reflect"
)

func main() {
	//修改基础类型
	var a int = 1024
	valueOfA := reflect.ValueOf(&a)
	valueOfA.Elem().SetInt(11019)
	fmt.Println(a) //11019

	type dog struct {
		LegCount int
	}

	valueOfDog := reflect.ValueOf(&dog{}).Elem()
	vLegCount := valueOfDog.FieldByName("LegCount")
	vLegCount.SetInt(1024) //1024
	fmt.Println(vLegCount)
}
