package main

import "fmt"

func main() {

	//武松带队行军 10 个小时
	//武松带队行军 20 个小时
	//武松带队行军 30 个小时
	//60
	//松江带队行军 1 个小时
	//松江带队行军 2 个小时
	//松江带队行军 4 个小时
	//7

	baseFunc1 := GetDoTaskFunc()
	baseFunc2 := GetDoTaskFunc()

	baseFunc1("武松", 10)
	baseFunc1("武松", 20)
	new1 := baseFunc1("武松", 30)
	fmt.Println(new1)

	baseFunc2("松江", 1)
	baseFunc2("松江", 2)
	new2 := baseFunc2("松江", 4)
	fmt.Println(new2)
}

//闭包的作用是保存【各自内层函数的状态，及下面progress】
func GetDoTaskFunc() func(name string, hours int) (progress int) {
	//progress只有函数在第一次调用时才会被初始化，之后都是调用上一次保存的值
	var progress int

	//定义匿名函数赋值给f
	f := func(name string, hours int) int {
		fmt.Printf("%s带队行军 %d 个小时\n", name, hours)
		progress += hours
		return progress
	}
	return f

}
