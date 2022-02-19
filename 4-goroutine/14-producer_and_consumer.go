package main

import "fmt"

func producerProcess(sendChan chan<- int) {
	for i := 0; i < 10; i++ {
		sendChan <- i
	}
	close(sendChan)
}

func consumer(receiveChan <-chan int) {
	for value := range receiveChan {
		fmt.Println("receive from：", value)
	}
}

func main() {
	channel := make(chan int)
	go producerProcess(channel)
	consumer(channel)
}
