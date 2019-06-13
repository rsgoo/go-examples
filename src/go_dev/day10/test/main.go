package main

import "fmt"

func sum(n1 int, n2 int) (int) {
	//defer 是先进后出的顺序
	defer fmt.Println("ok1 n1 = ", n1) //execute step 2
	//defer fmt.Println("ok2 n2 = ", n2) //execute step 1
	res := n1 + n2
	fmt.Println("sum is  ", res) //execute step 0
	return res
}

func main() {
	result := sum(11, 22)
	fmt.Println(result) //execute step 3
}
