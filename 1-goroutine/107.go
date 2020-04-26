package main

import "fmt"

func main() {
	f := func() bool {
		return false
	}

	//此时f();赋值结果被废弃，switch 默认匹配case 为 true的式子

	switch f(); {
	case true:
		//执行到这里
		fmt.Printf("true-")
	case false:
		fmt.Printf("false-")
	}

	fmt.Println("没有分号")

	switch f() {
	case true:
		fmt.Printf("true-")
	case false:
		//执行到这里
		fmt.Printf("false-")
	}

}
