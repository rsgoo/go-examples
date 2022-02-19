package main

import (
	"fmt"
	"strconv"
	"time"
)

//1: 生产者每一秒生产一个商品，并通知物流取货
//2: 物流将商品运输到商店
//3: 消费者阻塞等待从商店消费
//4: 消费10轮主协程结束

type Goods struct {
	Name string
}

var quitChan = make(chan interface{})

func main() {
	//生产者库房
	storageChan := make(chan Goods, 10)
	shopChan := make(chan Goods, 10)

	go Logistics(storageChan, shopChan)
	go Producer(storageChan)
	go Consumer(shopChan)

	for value := range quitChan {
		fmt.Println(value)
		close(shopChan)
		close(storageChan)
		close(quitChan)
	}
}

// Producer 生产者,只写
func Producer(storageChan chan<- Goods) {
	for {
		name := time.Now().Second()
		p := Goods{Name: "商品:" + strconv.Itoa(name)}
		fmt.Println("Producer ", p)
		storageChan <- p
		time.Sleep(time.Second)
	}
}

// Logistics 物流
// storageChan 只读chan
// shopChan 只写chan
func Logistics(storageChan <-chan Goods, shopChan chan<- Goods) {
	for {
		p := <-storageChan
		fmt.Println("transmit ", p)
		shopChan <- p
	}
}

// Consumer 消费者
func Consumer(shopChan <-chan Goods) {
	consumeCount := 0
	for {
		consumeCount++
		p := <-shopChan
		fmt.Println("consumer ", p)
		if consumeCount > 10 {
			break
		}
	}
	//用于通知主协程
	quitChan <- "go over"
}
