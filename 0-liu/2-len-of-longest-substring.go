package main

import "fmt"

//
func main() {
	//s := "pww12d"
	s := "pwwkew"
	fmt.Printf("%d", LenOfLongestSubstring(s))
}

func LenOfLongestSubstring(str string) int {
	locations := make([]int, 128)

	for i := 0; i < len(locations); i++ {
		locations[i] = -1
	}

	leftPos := 0
	strMaxLen := 0

	for i := 0; i < len(str); i++ {
		fmt.Println(locations[str[i]], str[i])
		if locations[str[i]] >= leftPos {
			leftPos = locations[str[i]] + 1
		}

		if i+1-leftPos >= strMaxLen {
			strMaxLen = i + 1 - leftPos
		}

		locations[str[i]] = i
	}

	return strMaxLen
}
