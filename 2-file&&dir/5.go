package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	fileName := "access.log";
	data := []byte("hello, welcome to use golang")
	err := ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		fmt.Println("err is ", err)
	}
}
