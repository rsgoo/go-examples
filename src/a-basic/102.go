package main

import "fmt"

func main() {
	var str = "hello,雨醉风尘"
	myStr := []int32(str)
	//for i := 0; i < len(myStr); i++ {
	//	fmt.Println(string(myStr[i]))
	//}

	myStr[0] = 'H'
	fmt.Println(string(myStr))
	for _, value := range myStr {
		fmt.Println(string(value))
	}

	//字符串的反转方法一
	var hello = "hello"
	var reverse = ""
	for i := len(hello) - 1; i >= 0; i-- {
		reverse += string(hello[i])
	}
	fmt.Println(reverse)

	//字符串的反转方法二
	myHello := []rune("123456")
	strLen := len(myHello)
	fmt.Println(strLen)

	for i := 0; i < strLen/2; i++ {
		myHello[i], myHello[strLen-1-i] = myHello[strLen-1-i], myHello[i]
	}

	fmt.Println(string(myHello))

}
