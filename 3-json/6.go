package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonStr = `[{"age":2,"gender":"Male","hobby":["watch tv","play mobile","chat with tom"],"name":"Jerry","rmb":268000},{"age":3,"gender":"Male","hobby":["watch tv","play mobile","chat with jerry"],"name":"Tom","rmb":288000}]`

	var jsonStrBytes = []byte(jsonStr)

	var slices = make([]map[string]interface{}, 0)

	err := json.Unmarshal(jsonStrBytes, &slices)
	if err != nil {
		fmt.Println("ers is ", err)
	}
	fmt.Println(slices)
}
