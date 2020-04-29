package main

import "fmt"

func main() {
	myChan := make(chan int, 5)

	myChan <- 111

	close(myChan)

	x, ok := <-myChan
	if ok {
		fmt.Println(x)
	} else {
		fmt.Println("管道读取失败")
	}

	fmt.Println("读取关闭后的管道：", <-myChan)
}
