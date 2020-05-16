package main

import (
	"fmt"
	"sort"
)

//287. 寻找重复数
//给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），
//可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

func main() {
	var nums = []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicateBySort(nums))
	fmt.Println(findDuplicateByMath(nums))
	fmt.Println(findDuplicateByMap(nums))
}

//方法一：排序后比较
func findDuplicateBySort(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return nums[i]
		}
	}
	return 0
}

//方法二：使用等差数列的思想，采用此方法有条件限制：切片元素必须是1到n之间的每一位
//1：算出 n个数(1-n)的和 Sn = (A1+An)*n/2,因为是n+1个数，所以最大数值 = 总个数-1
//2：算出n+1个数字的和()

func findDuplicateByMath(nums []int) int {
	numsLen := len(nums)
	maxVal := numsLen - 1

	commonSum := (1 + maxVal) * maxVal / 2
	numsSum := 0

	for i := 0; i < numsLen; i++ {
		numsSum += nums[i]
	}
	return numsSum - commonSum
}

//使用map方法
func findDuplicateByMap(nums []int) int {
	maps := make(map[int]bool)
	for _, val := range nums {
		if maps[val] {
			return val
		} else {
			maps[val] = true
		}
	}
	return -1
}
