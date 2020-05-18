package z_utils

func GetMaxVal(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func GetMinVal(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
