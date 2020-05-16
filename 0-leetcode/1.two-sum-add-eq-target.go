package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{3, -1, 6, 7, 1, 5, 8, 2, 10, 991, -112}
	var target = 9

	pos1 := twoSumWithForce(nums, target)
	fmt.Println(pos1)
	fmt.Println()

	pos2 := twoSumWithMap(nums, target)
	fmt.Println(pos2)
	fmt.Println()

	sortPos := twoSumTargetWithSort(nums, target)
	fmt.Println(sortPos)
}

//var nums = []int{3, -1, 6, 7, 1, 8, 2, 10, 991, -112}
//var target = 9
//针对排序的数组求两数之和为目标值
func twoSumTargetWithSort(arr []int, target int) [][]int {
	sort.Ints(arr)
	fmt.Println(arr)
	leftPos := 0
	rightPos := len(arr) - 1
	var posArr = [][]int{}
	for leftPos <= rightPos {
		if arr[rightPos]+arr[leftPos] == target {
			tmp := []int{leftPos, rightPos}
			posArr = append(posArr, tmp)
			leftPos++
			rightPos--
		} else if arr[rightPos]+arr[leftPos] < target {
			leftPos++
		} else {
			rightPos--
		}
	}
	return posArr
}

//暴力求解-无需排序，直接遍历
func twoSumWithForce(nums []int, target int) []int {
	targetSumIndexArr := []int{}
	for i := 0; i < len(nums); i++ {
		reqVal := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if reqVal == nums[j] {
				fmt.Printf("%d = %d + %d\n", target, nums[i], reqVal)
				targetSumIndexArr = append(targetSumIndexArr, i, j)
			}
		}
	}
	return targetSumIndexArr
}

//使用map
//range 后的
func twoSumWithMap(nums []int, target int) []int {
	result := []int{}
	maps := make(map[int]int)

	//{3, 6, 7, 1, 8, 2}
	for key, val := range nums {
		diffVal := target - val
		existVal, existFlag := maps[diffVal]
		if existFlag {
			result = append(result, key, existVal)
		}
		//maps赋值一定要在最后设置
		maps[val] = key
	}
	return result
}
