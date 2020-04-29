package main

import (
	"runtime"
	"fmt"
)

func main()  {
	//设置可用逻辑CPU核心书
	fmt.Println(runtime.GOMAXPROCS(1))
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS())
}
