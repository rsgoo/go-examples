package demo

func Jiechen(n int64) int64 {
	if n < 2 {
		return 1
	}
	return n * Jiechen(n-1)
}

func JS(n int64) int64 {
	var sum int64 = 0;
	var i int64 = 1;
	for ; i <= int64(n); i++ {
		sum += Jiechen(int64(i))
	}
	return sum
}
