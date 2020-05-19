package main

import "fmt"

//70. 爬楼梯
//leetcode-cn.com/problems/climbing-stairs/
func main() {
	//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	//每次你可以爬 1 或 2 或 3 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	for i := 1; i <= 4; i++ {
		fmt.Println(climbStairRecursive(i))
	}
	fmt.Println(climbStair(4))
}

//递归方式实现
func climbStairRecursive(n int) int {
	//爬一层楼梯：(1) == 1
	if n == 1 {
		return 1
	}
	//爬两层楼梯： (1+1) + (2) == 2
	if n == 2 {
		return 2
	}

	//爬三层楼梯： (1+1+1) + (2+1) + (1+2) + (3) == 4
	if n == 3 {
		return 4
	}
	return climbStairRecursive(n-1) + climbStairRecursive(n-2) + climbStairRecursive(n-3)
}

//迭代方式实现
func climbStair(n int) (rsp []int) {
	rsp = []int{1, 2, 4}

	for i := 3; i < n; i++ {
		rsp = append(rsp, rsp[i-3]+rsp[i-2]+rsp[i-1])
	}
	return rsp
}

// 爬4层楼梯
// 1 3(4) 【111,12,21,3】
// 2 2(2) 【11，2】
// 3 1(1) 【1】