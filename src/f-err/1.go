package main

import (
	"math"
	"errors"
	"fmt"
)

func main() {

	//defer是函数执行之前一定会执行
	defer func() {
		errReason := recover()
		if errReason != nil {
			fmt.Println("err exist。is ", errReason)
		}
	}()

	val, err := GetBallVolume(-10)
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	fmt.Println("小球的体积是：",val)

}

func GetBallVolume(radius float64) (vol float64, err error) {

	if radius < 0 {
		panic("半径不能为负数")
	}

	//如果半径不在取值范围内，温和的返回错误
	if radius < 5 || radius > 50 {
		err = errors.New("半径的取值返回是【5，50】")
		return
	}

	val := (4 / 3.0) * math.Pi * math.Pow(radius, 3.0)
	return val, nil
}
