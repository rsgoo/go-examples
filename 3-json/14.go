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

	srcJsonFile, err := os.OpenFile("animals.json", os.O_RDONLY, 0666)
	defer srcJsonFile.Close()

	if err != nil {
		fmt.Println("err is ", nil)
		return
	}

	jsonDecoder := json.NewDecoder(srcJsonFile)
	err1 := jsonDecoder.Decode(&slices)
	if err1 != nil {
		fmt.Println("err1 is ", err1)
		return
	}

	for _, value := range slices {
		fmt.Println(value)
	}
}
