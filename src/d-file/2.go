package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	file, err := os.OpenFile("1.txt", os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err is: ", err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		fmt.Println(reader.Text())
	}
}
