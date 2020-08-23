package main

import "fmt"

func main() {
	chan1 := make(chan int)
	fmt.Printf("%T\n", chan1)
	chan2 := make(chan bool)
	go func() {
		data, ok := <-chan1
		if ok {
			fmt.Println("读取到 chan1 的数据:", data)
		}
		chan2 <- true
	}()

	chan1 <- 11019
	//用于确定chan1已经被读取到
	<-chan2
	fmt.Println("main over")
}
