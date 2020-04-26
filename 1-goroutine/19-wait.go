package main

import (
	"fmt"
	"time"
	"sync"
)

type Human struct {
	Name string
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("协程A开始工作")
		time.Sleep(time.Second * 3)
		fmt.Println("协程A结束工作")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("协程B开始工作")
		timer := time.After(time.Second * 3)
		fmt.Println("协程B结束工作:", <-timer)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("协程C开始工作:")
		tiker := time.NewTicker(time.Second * 1)
		for i := 0; i < 4; i++ {
			<-tiker.C
		}
		tiker.Stop()
		fmt.Println("协程C结束工作:")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main over")
}
