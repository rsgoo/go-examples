package main

import (
	"fmt"
	"reflect"
)

type Enum int

const Zero Enum = 0

func main() {
	type cat struct {
	}

	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println("typeOfCat Type:", typeOfCat.Name()) //cat
	fmt.Println("typeOfCat Kind:", typeOfCat.Kind()) //struct

	TypeOfZero := reflect.TypeOf(Zero)
	fmt.Println("typeOfCat Type:", TypeOfZero.Name()) //Enum
	fmt.Println("typeOfCat Kind:", TypeOfZero.Kind()) //int

	typeOfCatPtr := reflect.TypeOf(&cat{})
	//指针变量的类型名称是空
	fmt.Println("typeOfNewCat Type:", typeOfCatPtr.Name()) //无值
	fmt.Println("typeOfNewCat Kind:", typeOfCatPtr.Kind()) //ptr

	insElem := typeOfCatPtr.Elem()
	fmt.Println("insElme Type:", insElem.Name())
	fmt.Println("insElme Kind:", insElem.Kind())
}
