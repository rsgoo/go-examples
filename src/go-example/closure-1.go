package main

import (
	"fmt"
)

func main() {

	//武松带队行军 10 个小时
	//武松带队行军 20 个小时
	//武松一共行军30小时
	//宋江带队行军 11 个小时
	//宋江带队行军 22 个小时
	//宋江一共行军33小时


	baseFunc1 := GetExecTaskFunc("武松")
	baseFunc2 := GetExecTaskFunc("宋江")

	baseFunc1(10)
	p := baseFunc1(20)
	fmt.Printf("武松一共行军%d小时\n", p)

	baseFunc2(11)
	p1 := baseFunc2(22)
	fmt.Printf("宋江一共行军%d小时\n", p1)

}

//闭包的作用是保存【各自内层函数的状态，及下面progress】
func GetExecTaskFunc(name string) func(hours int) (progress int) {
	var progress int
	//定义匿名函数赋值给f
	f := func(hours int) int {
		fmt.Printf("%s带队行军 %d 个小时\n", name, hours)
		progress += hours
		//匿名函数的返回类型是int(process)
		return progress
	}
	//GetExecTaskFunc 的返回类型是函数
	return f

}
