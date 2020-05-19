package main

import "fmt"

//最大公约数和最小公倍数获取
func main() {
	fmt.Println(getMinGBS(10, 15))
	fmt.Println(getMinGBSForce(10, 15))

	fmt.Println(getMinGBS(3, 6))
	fmt.Println(getMinGBSForce(3, 6))

	fmt.Println(getMinGBS(12, 14))
	fmt.Println(getMinGBSForce(12, 14))

	fmt.Println(getMinGBS(48, 18))
	fmt.Println(getMinGBSForce(48, 18))

	fmt.Println(getMaxGYS(10, 2))
	fmt.Println(getMaxGYSForce(10, 2))

	fmt.Println(getMaxGYS(60, 40))
	fmt.Println(getMaxGYSForce(60, 40))
}

//最小公倍数
func getMinGBS(a, b int) int {
	if a == 0 || b == 0 {
		return -1
	}

	var min int
	var max int
	if a >= b {
		min, max = b, a
	} else {
		min, max = a, b
	}

	modVal := max % min
	if modVal == 0 {
		return max
	} else {
		//获取最大公约数
		maxGYS := getMaxGYS(max, min)
		if (max%maxGYS == 0) && (min%maxGYS == 0) {
			//return modVal * (max / modVal) * (min / modVal)
			return max * min / maxGYS
		} else {
			return max * min
		}
	}
}

//暴力法求解最小公倍数
func getMinGBSForce(a, b int) int {
	if a == 0 || b == 0 {
		return -1
	}

	var min int
	var max int
	if a >= b {
		min, max = b, a
	} else {
		min, max = a, b
	}

	total := min * max

	for i := max; i <= total; i++ {
		if i%min == 0 && i%max == 0 {
			return i
		}
	}
	return total
}

//最大公约数
func getMaxGYS(a, b int) int {
	if a == 0 || b == 0 {
		return -1
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

//暴力获取最大公约数
func getMaxGYSForce(a, b int) int {
	if a == 0 || b == 0 {
		return -1
	}

	var min int
	var max int
	if a >= b {
		min, max = b, a
	} else {
		min, max = a, b
	}

	for i := min; i >= 1; i-- {
		if min%i == 0 && max%i == 0 {
			return i
		}
	}
	return 1
}
