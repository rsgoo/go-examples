package main

import (
	"fmt"
	"sync"
	"time"
)

//需求：计算1-200各数的阶乘，并把各个数的阶乘结果放置在map中
//最后显示出来，要求使用goroutine完成

//思路：
//1:编写一个函数，计算出阶乘，并访问map中
//2:启动多个协程

var MyMap = make(map[int]int64)

var Lock sync.Mutex

func Compute(n int) {
	base := 1
	for i := 1; i <= n; i++ {
		base *= i
	}
	Lock.Lock()
	MyMap[n] = int64(base)
	Lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go Compute(i)
	}
	time.Sleep(time.Second * 1)
	Lock.Lock()
	for key, value := range MyMap {
		fmt.Printf("map[%d] = %d\n", key, value)
	}
	Lock.Unlock()
}
