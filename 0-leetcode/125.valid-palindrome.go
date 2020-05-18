package main

import (
	"fmt"
	"strings"
)

//125. 验证回文串
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

func main() {
	str := "6V:.z;ID6DI;z.:VH"
	fmt.Println(isPalindromeStr(str))
}

func isPalindromeStr(s string) bool {
	s = strings.ToLower(s)
	bytes := []rune(s)
	//fmt.Println(bytes)
	var TraverseChar []rune
	TraverseChar = make([]rune, 0)
	//先过滤掉特殊字符
	for i := 0; i < len(bytes); i++ {
		if (bytes[i] >= '0' && bytes[i] <= '9') || (bytes[i] >= 'a' && bytes[i] <= 'z') {
			TraverseChar = append(TraverseChar, bytes[i])
		}
	}

	leftPos := 0
	rightPos := len(TraverseChar) - 1

	//然后使用双指针遍历
	for leftPos <= rightPos {
		if TraverseChar[leftPos] != TraverseChar[rightPos] {
			return false
		}
		leftPos++
		rightPos--
	}
	return true
}
