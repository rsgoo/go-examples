package main

import (
	"sdk"
	"fmt"
)

func main() {
	sdk.Greeting()
	fmt.Println(sdk.Sum(10, 20))
}
