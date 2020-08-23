package main

import "fmt"

//借助普通for循环和for ... range语句循环接收多个元素。
//遍历channel，遍历的结果就是接收到的数据，数据类型就是channel的数据类型。
//普通for循环接收channel数据，需要有break循环的条件；
//for … range会自动判断出channel已关闭，而无须通过判断来终止循环

func main() {
	ch1 := make(chan string)
	go sendData(ch1)

	//方式一：循环方式读取
	/*
		for {
			data := <-ch1
			if data == "" {
				fmt.Println("数据读取完毕")
				break
			}
			fmt.Println("get data - ", data)
		}
	*/

	//方式2
	/*
		for {
			data, ok := <-ch1
			if !ok {
				fmt.Println("数据读取完成")
				break
			}
			fmt.Println(data)
		}
	*/

	for data := range ch1 {
		fmt.Println(data)
	}

	fmt.Println("main func over")

}

func sendData(ch1 chan string) {
	defer close(ch1)
	for i := 0; i < 10; i++ {
		ch1 <- fmt.Sprintf("发送数据%d\n", i)
	}
	fmt.Println("数据发送完毕")
}
