package main

import (
	"errors"
	"fmt"
)

//<<<<<性别枚举
type Gender int

func (g Gender) String() string {
	return []string{"Male", "Female", "Other"}[g]
}

const (
	Male   = iota //0
	Female        //1
	Other         //2
)

//性别枚举>>>>>

type Human struct {
	Name          string
	Age           int
	Height        int
	Weight        int
	Looking       int //外貌
	TargetLooking int //目标外貌
	RMB           int
	Sex           Gender //自己的性别
	TargetSex     Gender
}

func (host *Human) Marry(other *Human) (happiness int, err error) {
	//性别不匹配
	if host.TargetSex != other.Sex {
		panic("emmmmmm")
	}

	//外表匹配
	if other.Looking < host.TargetLooking {
		err = errors.New("外貌不匹配")
		return
	}

	//计算幸福🥰程度
	happiness = (other.Height * other.RMB * other.Looking) / (other.Weight * other.Age)
	return

}

func NewHuman(name string, age, height, weight, looking, targetLooking, rmb int, sex, targetSex Gender) *Human {
	return &Human{
		Name:          name,
		Age:           age,
		Height:        height,
		Weight:        weight,
		Looking:       looking,
		TargetLooking: targetLooking,
		RMB:           rmb,
		Sex:           sex,
		TargetSex:     targetSex,
	}
}

func main() {
	andre := NewHuman("Andre", 25, 168, 120, 80, 80, 26800, Male, Female)
	anna := NewHuman("Andrena", 25, 165, 100, 70, 70, 25800, Female, Male)

	happiness, err := andre.Marry(anna)
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	fmt.Println(happiness)

}
