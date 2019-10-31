package main

import (
	"strconv"
	"time"
	"fmt"
)

type Product struct {
	Name string
}

func main() {
	//创建商店管道 和 计数管道
	shopChan := make(chan Product, 5)
	countChan := make(chan int, 5)
	go Produce(shopChan)
	go Consume(countChan, shopChan)

	for i := 0; i < 10; i++ {
		fmt.Println("从管道中读出：", <-countChan)
	}
	//close(shopChan)
	//close(countChan)

	fmt.Println("main over")
}


//生产中 shopChan 只写管道
func Produce(shopChan chan<- Product) {
	for {
		p := Product{"产品" + strconv.Itoa(time.Now().Second())}
		shopChan <- p
		fmt.Println("生产者生产：", p)
		time.Sleep(time.Second)
	}

}

//消费中 shopChan 只读管道
func Consume(countChan chan<- int, shopChan <-chan Product) {
	i := 0;
	for {
		x, ok := <-shopChan
		if ok {
			i++
			countChan <- i
			fmt.Println("消费者消费：", x)
			//消费10次
			//if i > 10 {
			//	//close(shopChan)
			//	//close(countChan)
			//	continue
			//}
		}
	}
}
