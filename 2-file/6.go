package main

import (
	"os"
	"fmt"
)

func main() {
	fileInfo, err := os.Stat("accessa.log")
	if err != nil {
		fmt.Println("err is ", err)
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}
		return
	}

	fmt.Println(fileInfo)
}
