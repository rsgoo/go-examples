package main

import "fmt"

func main() {

	baseArr := []int{9, 12, 34, 2, 24}
	//fmt.Println(bubbleSort(baseArr))
	//fmt.Println(quickSort(baseArr))
	//baseArr2 := []int{9, 12, 34, 2, 4}
	fmt.Println(quickSortV2(baseArr))

}

//冒泡排序
//冒泡排序是一种稳定排序算法。
func bubbleSort(arr []int) []int {
	arrLen := len(arr)

	if arrLen <= 1 {
		return arr
	}

	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//快速排序-使用外部空间
//快速排序不是一种稳定的排序算法。
func quickSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}

	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)
	baseVal := arr[0]

	for i := 1; i < arrLen; i++ {
		if baseVal > arr[i] {
			leftSlice = append(leftSlice, arr[i])
		} else {
			rightSlice = append(rightSlice, arr[i])
		}
	}

	leftSlice = quickSort(leftSlice)
	rightSlice = quickSort(rightSlice)

	return append(append(leftSlice, baseVal), rightSlice...)

}

//baseArr := []int{9, 12, 34, 2, 4}
func quickSortV2(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	l := 0
	r := len(data) - 1
	baseVal := data[0]

	for i := 1; i <= r; {
		fmt.Println(data)
		fmt.Println(baseVal, data[i])
		if baseVal < data[i] {
			fmt.Println("first")
			data[i], data[r] = data[r], data[i]
			r--
		} else {
			fmt.Println("second")
			//baseVal >= data[i]
			data[i], data[l] = data[l], data[i]
			l++
			i++
		}
		fmt.Println(data)
		fmt.Println()
	}
	quickSortV2(data[:l])
	quickSortV2(data[l+1:])

	return data
}
