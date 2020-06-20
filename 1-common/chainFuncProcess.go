package main

import (
	"fmt"
	"strings"
)

/**
 * @ 字符串的链式操作
 * @ strList 需要处理的字符串
 * @ funcList 需要链式调用的函数
 */
func StrProcess(strList []string, funcList []func(string) string) {
	for index, str := range strList {
		//str:处理过后需要保留的字符串
		for _, process := range funcList {
			str = process(str)
		}
		strList[index] = str
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	strList := []string{
		" go scanner ",
		" go parser",
		"go compiler",
		"go printer ",
		"go formatter ",
	}

	funcList := []func(string) string{
		removePrefix,
		strings.ToUpper,
		strings.TrimSpace,
	}

	StrProcess(strList, funcList)

	for _, str := range strList {
		fmt.Println(str)
	}
	//输出结果
	//GO SCANNER
	//GO PARSER
	//COMPILER
	//PRINTER
	//FORMATTER
}
