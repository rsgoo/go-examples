package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	goroutineNum := runtime.NumCPU()          //运行的协程数
	intChan := make(chan int, 1000)           //数字管道
	primeChan := make(chan int, 1000)         //素数管道
	exitChan := make(chan bool, goroutineNum) //退出标识

	//1：开启一个协程，向 intChan 放入 1-8000个数
	go putPrime(intChan)

	//2：开启四个协程，从 intChan 取出数据，并判断是否为素数，如果是，就放入到 primeChan 中
	for i := 0; i < goroutineNum; i++ {
		go calculatePrime(intChan, primeChan, exitChan)
	}
	fmt.Println(time.Now())
	fmt.Println(time.Since(now))

	go func() {
		for i := 0; i < goroutineNum; i++ {
			<-exitChan
		}
		//当我们从exitChan取出了4个退出标识，就可以关闭primeChan
		close(primeChan)
		close(exitChan)
	}()

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Println(val)
	}
	//fmt.Println(time.Since(now))
	fmt.Println("main over")
}

func putPrime(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}

func calculatePrime(intChan chan int, primeChan chan int, exitChan chan bool) {

	for {
		//是否是素数的标识：默认不是素数
		isPrime := true

		number, ok := <-intChan
		//读取不到管道数据时退出
		if !ok {
			break
		}

		//判断number是否是素数
		for i := 2; i < number; i++ {
			//不是素数
			if number%i == 0 {
				isPrime = false
				break
			}
		}

		//如果是素数，将改数字写入到素数管道
		if isPrime {
			primeChan <- number
		}
	}
	exitChan <- true
}
