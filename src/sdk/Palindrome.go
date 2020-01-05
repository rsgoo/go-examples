package sdk

//判断是否是回文字符串
func IsPalindrome(str string) bool {

	//默认是回文字符串
	palindromeFlag := true

	//fmt.Println(string(stringDetail))
	strDetail := []rune(str)

	//fmt.Println(len(strDetail))
	strLen := len(strDetail)
	for i := 0; i < strLen/2; i++ {
		//fmt.Println(i, ":", string(strDetail[i]))
		//fmt.Println(strLen-i-1, ":", strDetail[strLen-i-1])
		//fmt.Println("first：", strDetail[i])
		//fmt.Println("last：", strDetail[strLen-i-1])
		if strDetail[i] != strDetail[strLen-1-i] {
			palindromeFlag = false
		}
	}
	return palindromeFlag
}
