package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	Person := make(map[string]interface{})
	Person["name"] = "Jerry"
	Person["age"] = 2
	Person["rmb"] = 268000
	Person["gender"] = "Male"
	Person["hobby"] = []string{"watch tv", "play mobile", "chat with tom"}

	jsonFile, err := os.OpenFile("jerry.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println("file err is ", err)
	}

	//创建json编码器
	encoder := json.NewEncoder(jsonFile)

	//数据编码
	err1 := encoder.Encode(&Person)
	if err1 != nil {
		fmt.Println("err is ", err1)
		return
	}
}
