package main

import "fmt"

// iota 在const 关键字出现时将被重置为0。
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
// 使用iota能简化定义，在定义枚举时很有用。

const (
	a = iota //0
	b        //1
	_        //赋值为2，但是调过不显示
	d        //3
)

const (
	n1 = iota //0
	n2        //1
	n3 = 100  //100
	n4        //100，和上一个结果一样
	n5 = 200  //200
	n6        //200，和上一个结果一样
)

const (
	x1 = iota //0
	x2        //1
	x3 = 100  //100
	x4 = iota //3，
	x5        //4
	x6        //5，
	x7 = 200  //200
	x8 = iota //7，
	x9        //8，
)

const (
	_ = iota
	//N = (2 的 10次方)
	KB = 1 << (10 * iota) //iota = 1：1 * N
	MB = 1 << (10 * iota) //iota = 2：1 * N * N
	GB = 1 << (10 * iota) //iota = 3：1 * N * N * N
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	// A = O + 1; B = 0+2
	A, B = iota + 1, iota + 2 //1,2

	// C = A+1 = 2; D= B+1 = 3
	C, D

	// E = C+1 = 3; F= D+1 = 4
	E, F
)

func main() {
	fmt.Println("a:", a) //0
	fmt.Println("b:", b) //1
	fmt.Println("d:", d) //3

	fmt.Println("-----------")

	fmt.Println("n1 = :", n1)
	fmt.Println("n2 = :", n2)
	fmt.Println("n3 = :", n3)
	fmt.Println("n4 = :", n4)
	fmt.Println("n5 = :", n5)
	fmt.Println("n6 = :", n6)

	fmt.Println("-----------")

	fmt.Println("x1 = :", x1)
	fmt.Println("x2 = :", x2)
	fmt.Println("x3 = :", x3)
	fmt.Println("x4 = :", x4)
	fmt.Println("x5 = :", x5)
	fmt.Println("x6 = :", x6)
	fmt.Println("x7 = :", x7)
	fmt.Println("x8 = :", x8)
	fmt.Println("x9 = :", x9)

	//```````````````````````````````1 `2 `2 `3 `3 `4
	fmt.Println("A,B,C,D,E,F = ", A, B, C, D, E, F)
}
