package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	id  int64
	num int64
}

type Result struct {
	Item
	sum int64
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numChan := make(chan *Item, 10)
	resultChan := make(chan *Result, 10)
	exitChan := make(chan bool, 1)

	go producer(numChan)
	go customer(numChan, resultChan, exitChan)

	getPrint(resultChan)

	fmt.Println("main over")

}

func producer(numChan chan *Item) {

	var no int64
	for {
		number := rand.Int63n(100)
		item := &Item{
			id:  no,
			num: number,
		}

		numChan <- item
		no++
		fmt.Printf("== 序号%v : %v 生产完成 ==\n", no, number)
	}
	//close(numChan)
}

func customer(numChan chan *Item, resultChan chan *Result, exitChan chan bool) {
	for item := range numChan {
		result := &Result{
			Item: *item,
			sum:  item.num * 2,
		}
		resultChan <- result
	}
}

func getPrint(resultChan chan *Result) {
	for item := range resultChan {
		fmt.Printf("** 序号%v : 随机生产数 %v : sum = %v 消费完成 **\n", item.id, item.num, item.sum)
	}
}
