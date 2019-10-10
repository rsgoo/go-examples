package main

import (
	"io/ioutil"
	"fmt"
)

type Person struct {
} 

func main() {

	//file, err := os.OpenFile("1.txt", os.O_RDWR|os.O_APPEND, 0777)
	//if err != nil {
	//	fmt.Println("err is ", err)
	//}
	bytes, err1 := ioutil.ReadFile("1.txt")
	//xwbytes, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		fmt.Println(err1)
	}

	for _, value := range bytes {
		if len(string(value)) >0 {
			fmt.Printf(string(value))
		}
	}

	//fmt.Println(string(bytes))
}
