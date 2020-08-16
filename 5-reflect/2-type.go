package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Cat struct {
		Name string
		Type int `json:"type" id:"10"`
	}

	CatIns := Cat{
		Name: "Tom",
		Type: 1,
	}

	typeOfCat := reflect.TypeOf(CatIns)
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Println("name:", fieldType.Name)
		fmt.Println("type:", fieldType.Type)
		fmt.Println("-------------------")
	}

	catTypeName, ok := typeOfCat.FieldByName("Type")
	if ok {
		fmt.Println(catTypeName.Tag.Get("json"))
		fmt.Println(catTypeName.Tag.Get("id"))
	}

}
