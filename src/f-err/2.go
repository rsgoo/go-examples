package main

import (
	"math"
	"fmt"
)

//自定义异常 InvalidRadiusError
type InvalidRadiusError struct {
	Radius    float64
	MinRadius float64
	MaxRadius float64
}

//type error interface {
//	Error() string
//}

//这里的Error()实现error接口中的方法
func (err *InvalidRadiusError) Error() string {
	errInfo := fmt.Sprintf("%.2f是非法半径，合法的半径范围是[%.2f,%.2f]\n", err.Radius, err.MinRadius, err.MaxRadius)
	return errInfo
}

//工厂方法
func NewInvalidRadiusError(radius float64) *InvalidRadiusError {
	ire := &InvalidRadiusError{
		MinRadius: 5,
		MaxRadius: 50,
		Radius:    radius,
	}
	return ire

}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	val, err := GetBallVol(-2)
	if err != nil {
		fmt.Println("获取体积失败, err:", err)
		return
	}
	fmt.Println("体积是: ", val)
}

func GetBallVol(radius float64) (vol float64, err error) {

	if radius < 0 {
		//panic(&InvalidRadiusError{Radius: radius, MinRadius: 5, MaxRadius: 50})
		panic(NewInvalidRadiusError(radius))
	}

	//如果半径不在取值范围内，温和的返回错误
	if radius < 5 || radius > 50 {
		//err = &InvalidRadiusError{Radius: radius, MinRadius: 5, MaxRadius: 50}
		err = NewInvalidRadiusError(radius)
		return
	}

	val := (4 / 3.0) * math.Pi * math.Pow(radius, 3.0)
	return val, nil
}
