package main

import "fmt"

func main() {
	ch := make(chan int)

	go printer(ch)

	for i := 1; i <= 10; i++ {
		ch <- i
	}

	//标识写入完成
	ch <- 0 //A, 供X使用，通知写入器printer写入完成

	<-ch //B，阻塞，等待主协程退出标识
}

func printer(ch chan int) {

	for {
		data := <-ch
		if data == 0 { //X，判断依据来自A处管道写入数据
			break
		}
		fmt.Println(data)
	}
	ch <- 0 //Y，用于B处，通知主协程数据读取完成
}
