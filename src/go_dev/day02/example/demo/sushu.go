package demo

import "fmt"

func IsSushu(n int) string {
	var flag = true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			flag = false
			break
		}
	}
	var str string
	if flag {
		str = fmt.Sprintf("%d 是素数", n)
	} else {
		str = fmt.Sprintf("%d 不是素数", n)
	}
	return str;
}
