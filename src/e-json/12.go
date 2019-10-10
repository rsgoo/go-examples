package main

import (
	"os"
	"fmt"
	"encoding/json"
)

func main() {

	type Animal struct {
		Name   string
		Age    int
		Rmb    float64
		Gender string
		Hobby  []string
	}

	var slices = make([]Animal, 0)

	Tom := Animal{
		Name:   "TOM",
		Age:    2,
		Rmb:    26000,
		Gender: "male",
		Hobby:  []string{"eat fish", "play with jerry"},
	}

	Jerry := Animal{
		Name:   "TOM",
		Age:    2,
		Rmb:    26000,
		Gender: "male",
		Hobby:  []string{"eat fish", "play with tom"},
	}

	slices = append(slices, Tom)
	slices = append(slices, Jerry)

	jsonFile, err := os.OpenFile("animals.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	jsonEncoder := json.NewEncoder(jsonFile)
	err1 := jsonEncoder.Encode(slices)
	if err1 != nil {
		fmt.Println("err1 is ", err1)
		return

	}
}
