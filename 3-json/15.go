package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Animal struct {
		Name   string
		Age    int
		Rmb    float64
		Gender string
		Hobby  []string
	}
	var jsonStr = `{"age":2,"gender":"Male","hobby":["watch tv","play mobile","chat with tom"],"name":"Jerry","rmb":268000}`

	jsonBytes := []byte(jsonStr)

	var animal = &Animal{}

	err := json.Unmarshal(jsonBytes, animal)
	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	fmt.Println(*animal)
}
