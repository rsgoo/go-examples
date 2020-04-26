package main

import (
	"os"
	"fmt"
	"encoding/json"
)

func main() {

	Tom := make(map[string]interface{})

	srcJsonFile, err := os.OpenFile("jerry.json", os.O_RDONLY, 0666)
	defer srcJsonFile.Close()
	if err != nil {
		fmt.Println("err is ", nil)
		return
	}

	jsonDecoder := json.NewDecoder(srcJsonFile)
	err1 := jsonDecoder.Decode(&Tom)
	if err1 != nil {
		fmt.Println("err1 is ", err1)
		return
	}

	fmt.Println(Tom)
}
