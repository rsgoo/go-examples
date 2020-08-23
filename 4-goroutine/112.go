package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chan1 := make(chan int)
	ch_bool_1 := make(chan bool)
	ch_bool_2 := make(chan bool)
	ch_bool_3 := make(chan bool)

	rand.Seed(time.Now().UnixNano())

	//生产者
	go produce(chan1)

	go consume(1, chan1, ch_bool_1)
	go consume(2, chan1, ch_bool_2)
	go consume(3, chan1, ch_bool_3)

	<-ch_bool_1
	<-ch_bool_2
	<-ch_bool_3
}

func produce(chan1 chan int) {
	for i := 0; i <= 10; i++ {
		chan1 <- i
		fmt.Println("生产蛋糕，编号是：", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	defer close(chan1)
}

func consume(num int, chan1 chan int, ch chan bool) {
	for data := range chan1 {
		fmt.Printf("序号 %d 购买编号是%d 的蛋糕\n", num, data)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	ch <- true
	defer close(ch)

}
