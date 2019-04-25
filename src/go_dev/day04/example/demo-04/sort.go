package demo_04

import (
	"fmt"
	"sort"
)

//int 类型排序
func TestIntSort() {
	//不定长数组
	var a = [...]int{1, 34, 2, 908, 1832}
	b := a[:]
	fmt.Printf("%T\n", b)
	sort.Ints(b)
	fmt.Println(b)
}

//string 类型排序
func TestStringSort() {
	//不定长数组
	var a = [...]string{"php", "go", "rust", "PHP", "haskell"}
	b := a[:]
	fmt.Printf("%T\n", b)
	sort.Strings(b)
	fmt.Println(b)
}

func SearchInt() {
	var a = [...]int{1, 34, 2, 908, 1832,11019,1101}
	b := a[:]
	//sort.Ints(b)
	fmt.Println(b)
	index := sort.SearchInts(b, 11019)
	fmt.Println(index)
}
