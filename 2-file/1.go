package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	file, err := os.OpenFile("1.txt", os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err is: ", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	//创建文件的读取器
	for {
		myString, err1 := reader.ReadString('\n')
		if err1 == nil {
			fmt.Printf(myString)
		} else {

			if err1 == io.EOF {
				fmt.Println("文件读取完成 ", err1)
				break
			} else {
				fmt.Println("文件读取失败 err ", err1)
				break
			}
		}
	}

}
