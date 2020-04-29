package main

import (
	"time"
	"fmt"
)

func main() {
	chA := make(chan int, 5)

	chB := make(chan int, 4)
	chB <- 201
	chB <- 202
	chB <- 203
	chB <- 204

	chC := make(chan int, 3)
	chC <- 301
	chC <- 302
	chC <- 303

Outer:
	for {
		//select 随机选择一条能走通的case
		select {
		case chA <- 123:
			//执行任务A
			time.Sleep(time.Second * 1)
			fmt.Println("exec taskA - 写入", 123)
		case x := <-chB:
			//执行任务B
			time.Sleep(time.Second * 1)
			fmt.Println("exec taskB - ", x)
		case y := <-chC:
			//执行任务C
			time.Sleep(time.Second * 1)
			fmt.Println("exec taskC - ", y)
		default:
			fmt.Println("全部任务结束")
			//跳出select
			break Outer
		}
	}

	fmt.Println("main over")

}

//写
func TaskA(ch chan int) {
	for {
		fmt.Println("taskA")
		ch <- 123
		time.Sleep(time.Second * 1)
	}
}

//读
func TaskB(ch chan int) {
	for {
		fmt.Println("taskB")
		<-ch
		time.Sleep(time.Second * 1)
	}
}

//读
func TaskC(ch chan int) {
	for {
		fmt.Println("taskC")
		<-ch
		time.Sleep(time.Second * 1)
	}
}
