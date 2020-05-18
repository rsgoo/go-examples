package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "C:\\gitNote\\Go-tour\\src\\2-file\\1228.log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("openFile err: ", err)
	}
	ioutil.WriteFile()
}
