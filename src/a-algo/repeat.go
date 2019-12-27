package main

import "fmt"

//寻找一个集合中重复的数值
func main() {
	Method_1()
	Method_2()
}

// 使用双循环
// 时间复杂度 O(n*n)
func Method_1() {
	var arr []int
	arr = []int{3, 3, 1, 2, 5, 4, 7, 2, 4}
	fmt.Println(arr)
	sliceLen := len(arr)
	for i := 0; i < sliceLen-1; i++ {
		//fmt.Printf("arr[%d] = %d\n", i, arr[i])
		for j := i + 1; j < sliceLen; j++ {
			//fmt.Printf("arr[%d] = %d - ", j, arr[j])
			if arr[i] == arr[j] {
				fmt.Printf("《arr[%d],arr[%d] = %d, %d》\n", i, j, arr[i], arr[j])
			}
		}
	}
}

// 使用散列表方式
// 时间复杂度 O(n)
func Method_2() {
	var arr []int
	arr = []int{3, 1, 1, 2, 5, 4, 7, 2, 4}

	var arrMap map[int]int
	arrMap = make(map[int]int)

	fmt.Println(arr)
	for _, value := range arr {
		if arrMap[value] > 0 {
			arrMap[value] += 1
		} else {
			arrMap[value] = 1
		}
	}
	fmt.Println(arrMap)
}
