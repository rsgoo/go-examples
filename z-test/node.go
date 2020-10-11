package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel 1"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel 2"
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	go ping1(chan1)
	go ping2(chan2)

	select {
	case msg1 := <-chan1:
		fmt.Println(msg1)
	case msg2 := <-chan2:
		fmt.Println(msg2)
	case node := <-time.After(time.Millisecond * 2000):
		fmt.Println(node.Local().String(), "no message received. giving up")
	}

}
