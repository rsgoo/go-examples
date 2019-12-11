package main

import (
	"fmt"
	"time"
)

func main() {

	var grNum = 100

	var newNum = 10000000

	intChan := make(chan int, 2000)

	//素数chan
	primeChan := make(chan int, 2000)

	//退出标识管道
	exitChan := make(chan bool, grNum)
	var start = time.Now()
	//开启一个协程向intChan put 数据
	go putNum(newNum, intChan)

	//开启四个协程：从 intChan 读取数据，并判断是否为素数
	//如果是素数：放入primChan 管道
	for i := 0; i < grNum; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里对主线程进行处理
	go func() {
		for i := 0; i < grNum; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	//当从 exitChan 取出了四个结果，可以放心的关闭 primeNum

	//遍历 primeChan
	expire := time.Since(start).Seconds()
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数 = ", res)
	}

	fmt.Println("主线程退出,耗时：", expire)

}

//向 intChan 中放入 1-8000个数
func putNum(n int, intChan chan int) {
	for i := 1; i <= n; i++ {
		intChan <- i
	}
	//关闭 intChan 管道
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		//读不到管道数据的时候
		num, ok := <-intChan
		if !ok {
			break
		}

		//判断 num 是否是素数, 假定是素数
		primeFlag := true;
		for i := 2; i < num; i++ {
			//说明该num不是素数
			if num%i == 0 {
				primeFlag = false
				break
			}
		}

		//如果是素数
		if primeFlag {
			primeChan <- num
		}
	}

	//向 exitChan 管道写入 true
	exitChan <- true
}
