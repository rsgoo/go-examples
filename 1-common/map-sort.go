package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//随机种子
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu_%02d", i)
		val := rand.Intn(100)
		scoreMap[key] = val
	}

	scoreKeys := make([]string, 0, len(scoreMap))
	for key, _ := range scoreMap {
		scoreKeys = append(scoreKeys, key)
	}

	sort.Strings(scoreKeys)

	for _, key := range scoreKeys {
		fmt.Println(key, scoreMap[key])
	}

}
