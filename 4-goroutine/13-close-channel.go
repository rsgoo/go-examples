package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		for value := range ch {
			fmt.Println(value)
		}
		//读取协程一旦接收到管道关闭通知，就会结束阻塞读取
		fmt.Println("gorutine over")
	}()

	n := 0
	for {
		n++
		ch <- n

		if n > 10 {
			//协程关闭一定要放在break之前
			close(ch)
			break
		}
	}
	time.Sleep(time.Second * 10)
}
