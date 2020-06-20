package main

import "fmt"

func main() {
	baseArr := []int{9, 12, 34, 2, 4}
	fmt.Println(MergeSort(baseArr))
}

func MergeSort(arr []int) []int {
	count := len(arr)
	if count <= 1 {
		return arr
	}

	num := count / 2
	leftArr := MergeSort(arr[:num])
	rightArr := MergeSort(arr[num:])

	return SubSort(leftArr, rightArr)
}

func SubSort(leftArr, rightArr []int) (result []int) {
	leftPos := 0
	rightPos := 0

	for leftPos < len(leftArr) && rightPos < len(rightArr) {
		if leftArr[leftPos] < rightArr[rightPos] {
			result = append(result, leftArr[leftPos])
			leftPos++
		} else {
			result = append(result, rightArr[rightPos])
			rightPos++
		}
	}
	result = append(result, leftArr[leftPos:]...)
	result = append(result, rightArr[rightPos:]...)
	return
}
