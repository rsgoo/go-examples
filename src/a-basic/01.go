package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	go producer("Tom", channel)

	go producer("Jerry", channel)

	customer(channel)
}

func producer(name string, channel chan string) {
	//死循环，不停的生成数据
	for {
		channel <- fmt.Sprintf("%s:%v", name, rand.Int31())
		time.Sleep(time.Second)
	}

}

func customer(channel chan string) {
	//死循环，不停的生成数据
	for {
		message := <-channel
		fmt.Println(message)
	}

}
