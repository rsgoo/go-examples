package main

import (
	"strings"
	"fmt"
	"strconv"
	"time"
)

func main() {
	var str string = "hello, world; hello, dongdong";
	newStr := strings.Replace(str, "hello", "Hello", -1)
	fmt.Println(newStr)

	strCount := strings.Count(str, "o")
	fmt.Println(strCount)

	repeatStr := strings.Repeat("Rust\t", 3)
	fmt.Println(repeatStr)

	fieldStr := strings.Fields("haskell")
	fmt.Printf("%T\n", fieldStr)
	fmt.Println(fieldStr)

	splitStr := strings.Split("r-u-s-t", "-")
	fmt.Printf("%T\n", splitStr)
	fmt.Println(splitStr)

	fmt.Println(strconv.IntSize)

	t := time.Now()
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d-%02d-%4d\n", t.Hour(), t.Month(), t.Year())

}
