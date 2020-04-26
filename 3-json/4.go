package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	dataMap := make(map[string]interface{})

	var jsonStr = `{"age":2,"gender":"Male","hobby":["watch tv","play mobile","chat with tom"],"name":"Jerry","rmb":268000}`
	jsonStrBytes := []byte(jsonStr)
	err := json.Unmarshal(jsonStrBytes, &dataMap)
	if err != nil {
		fmt.Println("err is ", err)
	}
	fmt.Println(dataMap)
}
