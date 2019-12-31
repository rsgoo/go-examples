package main

import "fmt"

func writeData(intChan chan int) {
	for i := 1; i <= 500; i++ {
		intChan <- i
		fmt.Println("写入: ", i)
	}
	//写入完成后关闭管道
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		val, ok := <-intChan
		if !ok {
			break;
		}
		fmt.Println("读取: ", val)
	}

	//读取完成后，标识任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	//存放数据的管道
	intChan := make(chan int, 50)

	//存放退出标识的管道
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//根据能否读取到退出管道的元素判断 intChan 是否读取完成
	for {
		val, ok := <-exitChan
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
