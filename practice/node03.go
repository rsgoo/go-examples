package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	inputFile := "zero.log"
	outputFile := "zero_copy.log"
	bufRead, readErr := ioutil.ReadFile(inputFile)
	if readErr != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", readErr)
	}

	//将读取的文件内容存到变量中
	content := fmt.Sprintf("%s\n", bufRead)
	fmt.Println(content)

	writeErr := ioutil.WriteFile(outputFile, bufRead, 0777)
	if writeErr != nil {
		panic(writeErr.Error())
	}
}
