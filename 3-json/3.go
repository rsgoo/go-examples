package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var slices = make([]map[string]interface{}, 0)

	//var Jerry map[string]interface{}
	Jerry := make(map[string]interface{})
	Jerry["name"] = "Jerry"
	Jerry["age"] = 2
	Jerry["rmb"] = 268000
	Jerry["gender"] = "Male"
	Jerry["hobby"] = []string{"watch tv", "play mobile", "chat with tom"}

	slices = append(slices, Jerry)

	Tom := make(map[string]interface{})
	Tom["name"] = "Tom"
	Tom["age"] = 3
	Tom["rmb"] = 288000
	Tom["gender"] = "Male"
	Tom["hobby"] = []string{"watch tv", "play mobile", "chat with jerry"}

	slices = append(slices, Tom)

	bytes, err := json.Marshal(&slices)
	if err != nil {
		fmt.Println("err is ", err)
	}

	fmt.Println(string(bytes))
}
