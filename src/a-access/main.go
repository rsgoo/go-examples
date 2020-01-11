package main

import (
	"a-other"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main() {
	tom := &person.Person{
		Age:  proto.Int64(2),
		Name: proto.String("tom"),
		From: proto.String("tome"),
	}
	byteData, err := proto.Marshal(tom)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	newData := &person.Person{}
	err = proto.Unmarshal(byteData, newData)
	if err != nil {
		fmt.Println("err is :", err)
		return
	}

	fmt.Println(newData)
	fmt.Println(newData.GetAge())

}
