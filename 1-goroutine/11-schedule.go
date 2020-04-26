package main

import (
	"fmt"
)

func main() {

	//创建一个缓存容量为3的通知管道
	quitChan := make(chan string, 3)

	go Count("child-son", 10, quitChan)
	go Count("child-daughter", 10, quitChan)

	//主协程
	Count("Main", 5, quitChan)

	//不能在此处执行管道关闭操作，因为主协程执行完成后子协程并不一定能执行完成
	//close(quitChan)

	//阻塞等待从【任务完毕通知管道】读出所有完毕的消息
	for i := 0; i < 3; i++ {
		x := <-quitChan
		fmt.Println(x)
	}


	fmt.Println("=== main func over ===")
}

func Count(grName string, n int, quitChan chan string) {
	for i := 1; i <= n; i++ {
		fmt.Println(grName, i)
	}

	fmt.Println(grName, "工作完成")
	quitChan <- grName + " stask executed completed"
}
