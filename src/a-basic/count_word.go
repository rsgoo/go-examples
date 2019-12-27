package main

import (
	"strings"
	"fmt"
)

func main() {
	var str = "hello golang hello world hello rust"

	arr := strings.Split(str, " ")

	countMap := map[string]int{}

	for _, value := range arr {
		//判断arr 遍历出来的value 是否在 countMap 中存在
		_, ok := countMap[value]
		if ok {
			countMap[value] += 1
		} else {
			countMap[value] = 1
		}
	}

	fmt.Println(countMap)

}
