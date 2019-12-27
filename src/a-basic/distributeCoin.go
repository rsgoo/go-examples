package main

import (
	"strings"
	"fmt"
	"os"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew",
		"Sarah",
		"Augustus",
		"Heidi",
		"Emilie",
		"Peter",
		"Giana",
		"Adriano",
		"Aaron",
		"Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {

	userMap := make(map[string]int, len(users))

	//fmt.Println(strings.Fields("hello"))
	for _, userName := range users {
		transUserName := strings.ToUpper(userName)

		if strings.Index(transUserName, "E") >= 0 {
			fmt.Println(strings.Index(transUserName, "E"))
			fmt.Println(userName, " 执行1 ", )
			userMap[userName] += 1
		}

		if strings.Index(transUserName, "I") >= 0 {
			fmt.Println(strings.Index(transUserName, "I"))
			fmt.Println(userName, " 执行2 ", )
			userMap[userName] += 2
		}

		if strings.Index(transUserName, "O") >= 0 {
			fmt.Println(strings.Index(transUserName, "O"))
			fmt.Println(userName, " 执行3 ", )
			userMap[userName] += 3
		}

		if strings.Index(transUserName, "U") >= 0 {
			fmt.Println(strings.Index(transUserName, "U"))
			fmt.Println(userName, " 执行4 ", )
			userMap[userName] += 4
		}

		fmt.Println("------------------------------")
	}

	fmt.Println(userMap)

}
