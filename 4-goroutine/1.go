package main

import (
	"fmt"
)

func main() {
	go func(f func()) {
		for i := 0; i < 10; i++ {
			fmt.Println("协程打印:", i)
			//time.Sleep(time.Second * 1)
		}

		f()

	}(Greeting)

	for i := 10; i < 20; i++ {
		fmt.Println("主线程打印：", i)
		//time.Sleep(time.Second * 1)
	}

	//time.Sleep(time.Second * 1)
}

func Greeting()  {
	fmt.Println("Nice to meet you！")
}