package main

import "fmt"

func main() {
	var intChan = make(chan int, 5)
	var strChan = make(chan string, 5)

	for i := 1; i <= 5; i++ {
		intChan <- i
	}

	for i := 1; i <= 5; i++ {
		strChan <- fmt.Sprintf("hello, num %d", i)
	}

//Outer:
	for {
		//select 随机选择一条能走通的case
		select {
		case v := <-intChan:
			fmt.Println("receive from intChan: ", v)
		case v := <-strChan:
			fmt.Println("receive from intChan: ", v)
		default:
			//默认读取不到时的业务处理
			//fmt.Println("this is default value")
			//break Outer
			return
		}
	}
}
