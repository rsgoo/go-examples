package pressure

func PressureGetSum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func PressureGetRecursiveSum(n int) (sum int) {
	if n < 2 {
		return n
	}

	return n + PressureGetRecursiveSum(n-1)
}

func PressureGetFibonacciRecursive(n int) (sum int) {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return n
	}

	return PressureGetFibonacciRecursive(n-1) + PressureGetFibonacciRecursive(n-2)
}

func PressureGetFibonacci(n int) []int {
	slices := make([]int, n)
	slices[0], slices[1] = 1, 1

	for i := 2; i < n; i++ {
		slices[i] = slices[i-1] + slices[i-2]
	}
	return slices

}

func PressureGetFibonacciClosure() func() int {
	a1, a2 := 0, 1
	return func() int {
		a1, a2 = a2, a1+a2
		return a1
	}
}
