package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	//多读一写
	var rwm sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			rwm.RLock()
			fmt.Println("读取数据库 - ", index)
			<-time.After(time.Second * 3)
			rwm.RUnlock()
			wg.Done()
		}(i)

		wg.Add(1)
		go func(index int) {
			rwm.Lock()
			fmt.Println("写入数据库 - ", index)
			<-time.After(time.Second * 3)
			rwm.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("main over")
}
