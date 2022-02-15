package main

import "fmt"

func main() {
	deferCall()
}

func deferCall() {
	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")

}

//解析：defer的执行顺序是先进后出。当出现panic语句时，会按照defer先进后出的次序执行，然后执行panic
//output
/*
打印后
打印中
打印前
panic: 触发异常
goroutine 1 [running]:
main.deferCall()
        E:/Coding/golang/go-examples/days/1.go:22 +0x6b
main.main()
        E:/Coding/golang/go-examples/days/1.go:6 +0x17
exit status 2

*/
