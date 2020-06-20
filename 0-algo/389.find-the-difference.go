package main

import "fmt"

//389. 找不同
//给定两个字符串 s 和 t，它们只包含小写字母。
//字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//请找出在 t 中被添加的字母。

func main() {
	var str1 = "abcd"
	var str2 = "abcde"
	fmt.Println(string(findTheDifference(str1, str2)))
	fmt.Println(string(findTheDifferenceByHash(str1, str2)))
}

//len(s) + 1 = len(t)
func findTheDifference(s string, t string) byte {
	countS := 0
	for _, i2 := range s {
		countS += int(i2)
	}
	//fmt.Println(countS)

	countT := 0
	for _, i2 := range t {
		countT += int(i2)
	}
	//fmt.Println(countT)

	return byte(countT - countS)
}
