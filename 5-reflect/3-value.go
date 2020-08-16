package main

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string
	float32
	bool
	next *dummy
}

func main() {
	structIns := dummy{
		next: &dummy{},
	}

	d := reflect.ValueOf(structIns)
	fmt.Println("numField:", d.NumField())
	fmt.Println("field 2:", d.Field(2))
	fmt.Println("field 2 type:", d.Field(2).Type())

	fmt.Println("fieldByName(\"b\").type", d.FieldByName("b").Type())
	fmt.Println("fieldByIndex([]int{4,0})", d.FieldByIndex([]int{4, 0}).Type())

}
