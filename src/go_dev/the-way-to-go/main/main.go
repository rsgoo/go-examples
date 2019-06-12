package main

import (
	"fmt"
	"go_dev/the-way-to-go/way"
)

var a string

func main() {
	/*
	a = "G"   //全局
	print(a)  //G
	f1()
	print(a)	//R
	*/
	way.Rand()

}

func f1() {
	a := "O"   //局部
	print(a)   //O
	f2()
}

func f2() {
	a = "R"
	print(a)	//R
	f3()
}

func f3()  {
	fmt.Print(a)	//R
}