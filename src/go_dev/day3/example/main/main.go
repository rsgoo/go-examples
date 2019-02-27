package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a = 10
	var b *int
	b = &a
	*b = 11019
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println("***************")
	rand.Seed(time.Now().UnixNano())
	var randNum = rand.Intn(10)
	var input int
	fmt.Println("输入一个上个月")
	fmt.Scanln(&input)
	fmt.Println(randNum)
	fmt.Println(input)
}
