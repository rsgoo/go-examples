package main

import (
	"time"
	"fmt"
)

//周期性

func main() {
	//周期性定时器
	tikerStopFlag := false
	tiker := time.NewTicker(time.Second * 1)
	go func() {
		time.Sleep(time.Second * 5)
		tiker.Stop()
		tikerStopFlag = true

	}()

	for {
		if !tikerStopFlag {
			fmt.Println(<-tiker.C)
		} else {
			break
		}

	}
	fmt.Println("main over")
}
