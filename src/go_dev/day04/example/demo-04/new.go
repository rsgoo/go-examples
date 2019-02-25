package demo_04

import "fmt"

func NewCode() {
	var i int
	//0
	fmt.Println(i)

	j := new(int)
	//0xc00004a078
	fmt.Println(j)
	*j = 11019
	fmt.Println(*j)
	*j = 110100
	fmt.Println(*j)

	fmt.Println("this is about new and make difference")
	var s1 = new([]int)
	//&[]
	fmt.Println(s1)
	*s1 = make([]int, 1)
	(*s1)[0] = 1
	fmt.Println((*s1)[0])

	var s2 = make([]int, 10)
	//[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(s2)

	s2[0] = 11019
	//[11019 0 0 0 0 0 0 0 0 0]
	fmt.Println(s2)
}
