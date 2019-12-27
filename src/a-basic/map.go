package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var scoreMap = make(map[string]int, 100)

	for i := 1; i <= 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		val := rand.Intn(100)
		scoreMap[key] = val
	}

	var sortSlice = []string{}
	for key, _ := range scoreMap {
		sortSlice = append(sortSlice, key)
	}

	sort.Strings(sortSlice)

	for _, value := range sortSlice {
		fmt.Println(value, scoreMap[value])
	}

}
