package main

func main() {
	var myChan chan int

	//不能关闭没有初始化的管道
	//也不能关闭一个已经关闭的管道，即是不能重复关闭管道
	close(myChan)
}
