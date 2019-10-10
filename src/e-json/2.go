package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var Person map[string]interface{}
	Person = make(map[string]interface{})
	Person["name"] = "Jerry"
	Person["age"] = 2
	Person["rmb"] = 268000
	Person["gender"] = "Male"
	Person["hobby"] = []string{"watch tv", "play mobile", "chat with tom"}

	bytes, err := json.Marshal(Person)
	if err != nil {
		fmt.Println("err is ", nil)
	}
	fmt.Println(string(bytes))

}
