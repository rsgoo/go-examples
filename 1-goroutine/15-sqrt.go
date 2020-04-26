package main

import (
	"math"
	"strconv"
	"time"
	"fmt"
)

func GetSqrt(name string, n int, semaphoreChan chan string) {

	//要执行，先注册
	//能写入就执行，写不进去就阻塞到能写入为止
	semaphoreChan <- name

	ret := math.Sqrt(float64(n))
	time.Sleep(time.Second)
	fmt.Printf("%d的平方根是%.3f\n", n, ret)

	//任务执行完毕，从信号量控制管道注销自己，以便为其他并发协程腾出空间
	<-semaphoreChan
}

func main() {

	//并发数控制管道
	semaphoreChan := make(chan string, 5)	//最大并发数控制为5
	//begin := time.Now()
	//业务代码

	for i := 0; i <= 100; i++ {
		go GetSqrt("协程"+strconv.Itoa(i), i, semaphoreChan)
	}
	//spendTime := time.Since(begin)  //花费的时间
	//fmt.Println(spendTime)
	for {
		time.Sleep(time.Second)
	}
}
