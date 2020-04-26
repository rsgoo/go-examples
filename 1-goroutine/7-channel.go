package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	myChan := make(chan string, 3)

	//开启协程读取channel
	go func() {
		for i := 0; i < 5; i++ {
			//管道中没有数据时，产生读出阻塞
			x := <-myChan
			fmt.Println("读出channel: ", x)
		}
	}()

	//主协程写入管道， 最多智能时8
	for i := 0; i < 8; i++ {
		myChan <- strconv.Itoa(i)
		fmt.Println("已经写入", strconv.Itoa(i))
		time.Sleep(time.Millisecond * 10)
	}

}
