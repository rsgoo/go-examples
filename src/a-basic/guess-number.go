// code by shaoyongyang
package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	answer := myRand.Intn(1000)
	fmt.Println(rand.Intn(10))
	fmt.Println(answer)

	//for true {
	//	fmt.Println("请输入你猜想的数值:")
	//	var guess string
	//	fmt.Scan(&guess)
	//	guessNum, _ := strconv.Atoi(guess)
	//	if guessNum == answer {
	//		fmt.Println("beego 猜对了")
	//		break;
	//	} else if guessNum < answer {
	//		fmt.Println("猜小了")
	//	} else {
	//		fmt.Println("猜大了")
	//	}
	//}
	//fmt.Println("正确答案是:", answer)
}
