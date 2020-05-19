package main

import "fmt"

//41. 缺失的第一个正数
//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
func main() {
	nums := []int{1,2,0}
	//nums := []int{7, 8, 9, 11, 12}
	result := firstMissingPositive(nums)
	fmt.Println(result)
}

func firstMissingPositive(nums []int) int {
	numsCnt := len(nums)
	//+2是考虑到第一个元素为0和切片的maxVal加1
	mySlice := make([]int, numsCnt+2)
	fmt.Println(mySlice)

	for _, v := range nums {
		if v > 0 && v < numsCnt+2 {
			mySlice[v] = 1
		}
	}
	fmt.Println(mySlice)

	//因为第一位是0不参与判断
	for i := 1; i < len(mySlice); i++ {
		if mySlice[i] == 0 {
			return i
		}
	}
	return -1
}
//nums := []int{1,2,0}
//[0 0 0 0 0]
//[0 1 1 0 0]
//3