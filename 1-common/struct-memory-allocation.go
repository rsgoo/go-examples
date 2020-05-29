package main

import "fmt"

type test struct {
	a int
	b int
	c int
	d int
	e int
	f string
	g bool
}

func main() {
	node := test{1, 2, 3, 4, 5, "go", true}

	fmt.Printf("node   %p\n", &node)
	fmt.Printf("node.a %p\n", &node.a)
	fmt.Printf("node.b %p\n", &node.b)
	fmt.Printf("node.c %p\n", &node.c)
	fmt.Printf("node.d %p\n", &node.d)
	fmt.Printf("node.e %p\n", &node.e)
	fmt.Printf("node.f %p\n", &node.f)
	fmt.Printf("node.g %p\n", &node.g)
	//node   0xc000036040 ==> 4 * 16 + 0 = 64
	//node.a 0xc000036040 ==> 4 * 16 + 0 = 64
	//node.b 0xc000036048 ==> 4 * 16 + 8 = 72
	//node.c 0xc000036050 ==> 5 * 16 + 0 = 80
	//node.d 0xc000036058 ==> 5 * 16 + 8 = 88
	//node.e 0xc000036060 ==> 6 * 16 + 0 = 96
	//node.f 0xc000036068 ==> 6 * 16 + 8 = 104
	//node.g 0xc000036078 ==> 7 * 16 + 8 = 120
}
