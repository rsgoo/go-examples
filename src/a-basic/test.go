package main

import "fmt"

func main() {
	var a = []int{1}
	var b = a[:]

	//a 和 b 指向同一个内存地址
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b)

	for i := 2; i < 5; i++ {
		b = append(b, i)
		fmt.Printf("%v\n", b)
		fmt.Printf("len(b) = %d, cap(b) = %d \n", len(b), cap(b))
		fmt.Printf("add = %p\n", b)
		fmt.Println()
	}
	var x []int

	x = append(x, b...)
	fmt.Println(x)

	var n = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		n = append(n, i)
	}
	fmt.Println(n)
	fmt.Println(len(n), cap(n))

}
