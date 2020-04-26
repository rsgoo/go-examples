package main

import (
	"os"
	"fmt"
	"bufio"
	"time"
	"strconv"
)

func main() {
	file, err := os.OpenFile("access.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 07777)
	defer file.Close()
	if err != nil {
		fmt.Println("err is: ", err)
	}

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		y, m, d := time.Now().Date()
		str := fmt.Sprintf("%v-", y) +
			fmt.Sprintf("%v-", m) +
			fmt.Sprintf("%v ", d) +
			fmt.Sprintf("%v:", time.Now().Hour()) +
			fmt.Sprintf("%v:", time.Now().Minute()) +
			fmt.Sprintf("%v ", time.Now().Second())
		_, err1 := writer.WriteString(str + " hello" + strconv.Itoa(i) + "\n")
		if err1 != nil {
			fmt.Println("write err is ", err1)
		}
	}
	writer.Flush()
}
