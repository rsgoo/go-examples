package main

import (
	"fmt"
	"sync"
	"time"
)

//并发不安全案例
func unSafe() {
	var wg sync.WaitGroup
	var money = 2000
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				money += 1
			}
			wg.Done()
		}()
	}
	wg.Wait()
	//money 大概率小于2000+10*100
	fmt.Println(money)
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	var money = 2000
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(index int) {
			fmt.Printf("协程%d开始抢锁\n", index)
			//开始抢锁
			mu.Lock()
			fmt.Printf("协程%d抢锁成功，开始操作数据\n", index)
			for j := 1; j <= 1000; j++ {
				money += 1
			}
			<-time.After(time.Second * 5)
			fmt.Printf("协程%d将锁释放\n", index)
			//释放锁
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	//money 大概率小于2000+10*100
	fmt.Println(money)
}

//协程1开始抢锁
//协程3开始抢锁
//协程10开始抢锁
//协程5开始抢锁
//协程4开始抢锁
//协程7开始抢锁
//协程8开始抢锁
//协程9开始抢锁
//协程2开始抢锁
//协程1抢锁成功，开始操作数据
//协程6开始抢锁
