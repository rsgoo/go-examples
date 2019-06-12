package demo_04

import "fmt"

func Bubble(arr []int) {
	//fmt.Println("排序前：", *Array)
	//数组五个元素 两两比较需要4次
	for i := 0; i < len(arr)-1; i++ {
		//-i 是因为每一次排序后已经确定了一个最值
		for j := 0; j < len(arr)-1-i; j++ {
			if (arr)[j] > (arr)[j+1] {
				(arr)[j+1], (arr)[j] = (arr)[j], (arr)[j+1]
			}
		}
	}
	fmt.Println(arr)
	//fmt.Println("第一轮比较完后：", *Array)
}

