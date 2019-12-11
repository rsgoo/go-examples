package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 5; i >= 0; i-- {
			ch <- i * 2
		}
	}()

	for data := range ch {
		fmt.Println(data)

		//这里需要判断退出条件。否则如果继续发送，由于接收 goroutine 已经退出，没有 goroutine 发送到通道，因此 runtime(运行时) 会触发宕机报错。
		if data == 0 {
			break
		}
	}
}
