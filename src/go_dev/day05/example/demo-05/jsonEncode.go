package demo_05

import (
	"encoding/json"
	"fmt"
)

type StudentInfo struct {
	//
	Name  string `json:"name_1"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func JsonEncode() {
	var stu = StudentInfo{
		Name:  "stu1",
		Age:   18,
		Score: 80,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode failed, err：", err)
		return
	}

	jsonData := fmt.Sprintf("%s", string(data))
	fmt.Println(jsonData)

}
