package main

import (
	"fmt"
	"time"
)

//向 intChan 中放入 1-8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 10000; i++ {
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
		isPrime := true;
		for i := 2; i < num; i++ {
			//说明该num不是素数
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		//如果是素数
		if isPrime {
			primeChan <- num
		}
	}

	//向 exitChan 管道写入 true
	fmt.Println("跑完一次")
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)

	//素数chan
	primeChan := make(chan int, 2000)

	//退出标识管道
	exitChan := make(chan bool, 4)

	//开启一个协程向intChan put 数据
	begin := time.Now()
	go putNum(intChan)

	//开启四个协程：从 intChan 读取数据，并判断是否为素数
	//如果是素数：放入primChan 管道
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里对主线程进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(exitChan)
		close(primeChan)
	}()
	elapsed := time.Since(begin)
	//当从 exitChan 取出了四个结果，可以放心的关闭 primeNum

	//遍历 primeChan
	index := 1
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数 %v= %v\n", index, res)
		index++
	}

	fmt.Println("主线程退出")
	fmt.Println("app elapsed:", elapsed)

}
