package main

import "fmt"

func main() {
	var arr = []int{2, 3, 6, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}

func twoSum(numbers []int, target int) [][]int {
	result := [][]int{}
	left := 0
	right := len(numbers) - 1

	for left <= right {
		if numbers[left]+numbers[right] == target {
			temp := []int{left, right}
			result = append(result, temp)
			left++
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return result
}
