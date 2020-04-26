package main

import (
	"sync"
	"fmt"
	"time"
)

//创建40条协程，其中30条模拟读，10条模拟写
//最大并发数控制在5
//数据库允许一写多读
//使用定时器和秒表模拟IO耗时
//主协程必须恰好结束在所有协程完成工作的时候

func ReadDB(wg *sync.WaitGroup, ch chan int, rmMutex *sync.RWMutex) {
	ch <- 123
	rmMutex.RLock()
	fmt.Println("readDB")
	<-time.After(time.Second * 1)
	rmMutex.RUnlock()
	<-ch
	wg.Done()
}

func WriteDB(wg *sync.WaitGroup, ch chan int, rmMutex *sync.RWMutex) {
	ch <- 123
	rmMutex.Lock()
	fmt.Println("writeDB")
	ticker := time.NewTicker(1 * time.Second)
	for i := 0; i < 3; i++ {
		<-ticker.C
	}
	rmMutex.Unlock()
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var rmMutex sync.RWMutex

	//并发数
	chSemaphore := make(chan int, 10)

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			wg.Add(1)
			go ReadDB(&wg, chSemaphore, &rmMutex)
		}

	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WriteDB(&wg, chSemaphore, &rmMutex)
	}

	wg.Wait()

	fmt.Println("main over")
}
