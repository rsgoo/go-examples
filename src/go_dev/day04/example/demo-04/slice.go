package demo_04

import "fmt"

type slice struct {
	//切片的三个要素
	ptr *[100]int   //指针：指向数组
	len int
	cap int
}

func modify(s slice) {
	s.ptr[1] = 1000
}

func makeSlice(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func SliceCode() {
	var a []int
	a = append(a, 1, 11, 111)
	fmt.Println(a)
	a = append(a, a...)
	fmt.Println("数组 append 一个数组 使用后...", a)
}

func TestSlice() {
	var s1 slice
	s1 = makeSlice(s1, 10)
	s1.ptr[0] = 100
	modify(s1)
	fmt.Println(s1.ptr)

}
