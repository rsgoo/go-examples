package main

import (
	"time"
	"fmt"
)

func sayHello() {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test() {
	//使用错误处理机制 defer + recover
	defer func() {
		//捕获test 抛出的 panic
		err := recover()
		if err != nil {
			fmt.Println("test() 发送错误 ", err)
		}
	}()

	//定义一个map
	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok = ", i)
		time.Sleep(time.Second)
	}
}
