package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int64 = 7
	//进制转换
	fmt.Printf("%v", strconv.FormatInt(a, 2))
}
