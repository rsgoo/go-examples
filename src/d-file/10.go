package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	bytes, err := ioutil.ReadFile("access.log")
	if err != nil {
		fmt.Println("err is", err)
	}

	contentStr := string(bytes)

	var numCount = 0
	var letterCount = 0
	var spaceCount = 0
	for _, char := range contentStr {
		if char >= '0' && char <= '9' {
			numCount++
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			letterCount++
		} else if char == '\n' || char == '\r' || char == '\t' {
			spaceCount++
		}
	}
	fmt.Println("numCount is ", numCount)
	fmt.Println("letterCount is ", letterCount)
	fmt.Println("spaceCount is ", spaceCount)
}
