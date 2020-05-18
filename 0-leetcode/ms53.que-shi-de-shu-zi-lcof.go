package main

import "fmt"

func main() {
	nums := []int{0}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	maps := make([]int, len(nums)+1)
	for _, val := range nums {
		if val >= 0 && val < len(nums)+1 {
			maps[val] = 1
		}
	}
	fmt.Println(maps)
	for i := 0; i < len(maps); i++ {
		if maps[i] == 0 {
			return i
		}
	}
	return 0
}
