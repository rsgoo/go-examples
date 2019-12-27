package main

import "fmt"

func main() {
	var a = []int{}
	var b = a[:]

	//a 和 b 指向同一个内存地址
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("len(b) = %d, cap(b) = %d \n", len(b), cap(b))
	fmt.Printf("%v\n", b)

	for i := 0; i <= 5; i++ {
		b = append(b, i)
		fmt.Printf("%v\n", b)
		fmt.Printf("len(b) = %d, cap(b) = %d \n", len(b), cap(b))
		fmt.Printf("add = %p\n", b)
		fmt.Println()
	}

	var c = []int{1, 2, 3}
	fmt.Println(c)
	fmt.Println(len(c), cap(c))

	d := append(c, 4)
	fmt.Println(d)
	fmt.Println(len(d), cap(d))

	e := append(d, 5, 6, 7, 8, 9, 11, 12, 13, 14,15,16,17,18)
	fmt.Println(e)
	fmt.Println(len(e), cap(e))
}
