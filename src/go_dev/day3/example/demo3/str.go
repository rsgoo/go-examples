package demo3

import (
	"strings"
	"fmt"
	"strconv"
	"time"
)

func Prefix(source, prefix string) bool {
	return strings.HasPrefix(source,prefix)
}

func Test()  {
	flag := Prefix("https://github.com/inscode/Rust", "http")
	fmt.Printf("%v\n", flag)

	fmt.Println(strings.Index("hello,go", "g"))
	fmt.Println(strings.IndexByte("hello,go", 'o'))

	fmt.Println(strings.Replace("go hello,go", "go", "golang", -1))

	fmt.Println(strings.Fields("hello"))

	fmt.Println(strconv.Itoa(11019))

	//时间格式化
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
	start := time.Now().Second()
	fmt.Println(start)
	time.Sleep(time.Second * 5)
	end := time.Now().Second()
	fmt.Println(end)
	fmt.Printf("%d", end-start)
}
