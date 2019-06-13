package demo_04

import "fmt"

var a string

func Score() {
	a = "G"   //全局
	print(a)  //G
	f1()
	print(a)	//R
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