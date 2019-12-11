package main

import (
	"fmt"
	"g-book/chapter/goinstall/mypkg"
)

func main() {
	mypkg.CustomPkgFunc()
	fmt.Println("main func")
}
