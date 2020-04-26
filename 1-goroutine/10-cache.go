package main

import "fmt"

func main() {

	myChan := make(chan int, 5)
	myChan <- 101

	fmt.Println("管道的长度是：", len(myChan))
	fmt.Println("管道的容量是：", cap(myChan))

	myChan <- 102
	myChan <- 103
	myChan <- 104
	myChan <- 105
	fmt.Println("管道的长度是：", len(myChan))
	fmt.Println("管道的容量是：", cap(myChan))

	<- myChan
	<- myChan
	<- myChan
	<- myChan
	<- myChan
	fmt.Println("管道的长度是：", len(myChan))
	fmt.Println("管道的容量是：", cap(myChan))


	fmt.Println(<- myChan)

	fmt.Println("管道的长度是：", len(myChan))
	fmt.Println("管道的容量是：", cap(myChan))

}
