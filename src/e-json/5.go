package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name   string
	Age    int
	Rmb    float64
	Gender string
	Hobby  []string
}

func main() {

	var Tom = Cat{}
	var jsonStr = `{"Name":"Tom","Age":2,"Rmb":268000,"Gender":"Male","Hobby":["eat fish","play with tom"]}`

	jsonStrBytes := []byte(jsonStr)

	err := json.Unmarshal(jsonStrBytes, &Tom)
	if err != nil {
		fmt.Println("err is ", err)
	}
	fmt.Println(Tom)
}
