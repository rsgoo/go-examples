package main

import (
	"time"
	"fmt"
	"sync"
)

var myMap = make(map[int32]int32, 10)
var lock sync.Mutex

func test(n int32) {

	var i int32
	var res int32 = 1
	for i = 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {

	//开启多个协程完成任务
	var i int32
	for i = 1; i <= 10; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	//输出结果
	lock.Lock()
	for key, value := range myMap {
		fmt.Println(key, value)
	}
	lock.Unlock()
}
