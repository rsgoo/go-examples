package main

import (
	"time"
	"fmt"
	"os"
)

func main() {
	//AfterThreeSecond2()
	//TimerStop()
	TimerReset()
}

func AfterThreeSecond1() {
	//timer定时器
	//2019-11-02 19:31:15.975325
	fmt.Println(time.Now())
	timer := time.NewTimer(time.Second * 3)
	//2019-11-02 19:31:18.978344
	times := timer.C
	fmt.Println(<-times)
	fmt.Println("A Big News")
}

func AfterThreeSecond2() {
	fmt.Println("计时开始:", time.Now())
	times := <-time.After(time.Second * 3)
	fmt.Println("计时结束:", times)
	fmt.Println("A Big News")
}

//定时器终止
func TimerStop() {
	fmt.Println("开始时间：", time.Now())
	timer := time.NewTimer(time.Second * 5)

	go func() {
		time.Sleep(time.Second * 3)
		//定时器终止，不再向管道写入数据
		ok := timer.Stop()
		if ok {
			fmt.Println("拆除成功")
			os.Exit(0)
		}
	}()

	//阻塞5秒
	endTime := <-timer.C
	fmt.Println("引爆时间：", endTime)

}

//定时器重置
func TimerReset() {
	//开始时间： 2019-11-02 22:19:46.861756
	fmt.Println("开始时间：", time.Now())

	//五秒过后： 2019-11-02 22:19:51.863224
	timer := time.NewTimer(time.Second * 5)
	fmt.Println("五秒过后：", <-timer.C)

	//重置十秒： 2019-11-02 22:20:01.867238
	timer.Reset(time.Second * 10)
	fmt.Println("重置十秒：",  <-timer.C)
}
