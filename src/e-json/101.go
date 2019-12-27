package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Score int
	Grade string
}

type Student struct {
	Id     int      `json:"id"`
	Gender string   `json:"gender"`
	Name   string   `json:"name"`
	Detail *Student `json:"detail"`
}

func main() {
	var stu1 = Student{
		Id:     1,
		Gender: "男",
		Name:   "user1",
		Detail: &Student{
			Id:     11,
			Gender: "男男",
			Name:   "user2",
		},
	}

	jsonStr, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))

	var decode = Student{}
	json.Unmarshal([]byte(jsonStr), &decode)
	fmt.Println(decode.Name)

}
