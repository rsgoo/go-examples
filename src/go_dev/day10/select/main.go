package main

import (
	"strconv"
	"fmt"
)

func main() {
	//定义一个管道 10个数据 int
	intChan := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 1; i <= 5; i++ {
		str := "hello " + strconv.Itoa(i)
		stringChan <- str
	}

	//问题：在实际开发中，可能我们不好确定什么时候关闭该管道
	//可以使用select 方式解决
	for {
		select {
		case v := <-intChan:
			fmt.Println("从 intChan 读取数据 ", v)
		case v := <-stringChan:
			fmt.Println("从 strChan 读取数据 ", v)
		default:
			fmt.Println(":)")
			return	//使用return退出

		}
	}
assssssss}
