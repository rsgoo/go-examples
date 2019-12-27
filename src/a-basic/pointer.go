package main

import "fmt"

func main() {
	var a *int
	var b *string

	//<nil> : 0x0
	fmt.Printf("%v : %p\n", a, a)
	//<nil> : 0x0
	fmt.Printf("%v : %p\n", b, b)

	//使用new会分配一块内存空间，并将内存空间地址赋值给x,y
	//x,y 指向的值是new(Type)的零值
	//既是 *x = 0; *y = ""

	var x = new(int)
	var y = new(string)
	//0 : 0xc00004a090 : 0xc00004a090
	fmt.Printf("%v : %v : %p\n", *x, x, x)

	// : 0xc00003e1c0 : 0xc00003e1c0
	fmt.Printf("%v : %v : %p\n", *y, y, y)

	var c = new([3]int)
	(*c)[0] = 1
	(*c)[1] = 2
	(*c)[2] = 3
	fmt.Println(*c)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err is :", err)
		}
	}()

	var arr []int
	arr = make([]int, 3)
	fmt.Println(arr)
	fmt.Println(len(arr),cap(arr))
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
		fmt.Println(len(arr),cap(arr))
	}
	fmt.Println(arr)

}
