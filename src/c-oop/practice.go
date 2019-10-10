package main

import (
	"fmt"
	"math/rand"
	"time"
)

type AnimalInc interface {
	Death()
	Live()
}

//鸟
type Bird struct {
}

func (b *Bird) Live() {
	fmt.Printf("一只鸟儿在唱歌\n")
}

func (b *Bird) Death() {
	fmt.Printf("一只鸟儿在坠落\n")
}

//鱼
type Fish struct {
}

func (f *Fish) Live() {
	fmt.Printf("一只鱼在畅游\n")
}

func (f *Fish) Death() {
	fmt.Printf("红烧鱼\n")
}

//野兽
type Monster struct {
}

func (m *Monster) Live() {
	fmt.Printf("一只怪兽在唱歌\n")
}

func (m *Monster) Death() {
	fmt.Printf("一只怪兽在坠落\n")
}

func (m *Monster) Run() {
	fmt.Printf("一只怪兽在奔跑\n")
}

func (m *Monster) Hunt() {
	fmt.Printf("一只怪兽在捕食\n")
}

type BigTiger struct {
	Monster
	Name string
}

func (t *BigTiger) Hunt() {
	fmt.Println("老虎要捕食了")
}

type Human struct {
	Monster
	Name string
}

func (h *Human) Live() {
	fmt.Println("好好工作，努力挣钱，吃大餐")
}

func (h *Human) Hunt() {
	fmt.Println("周末了，好好吃一顿吧")
}

func main() {
	bird := Bird{}
	fish := Fish{}
	tiger := BigTiger{
		Monster: Monster{},
		Name:    "这是一只老虎",
	}
	human := Human{
		Monster: Monster{},
		Name:    "南粤食神",
	}

	animals := make([]AnimalInc, 0)
	animals = append(animals, &bird)
	animals = append(animals, &fish)
	animals = append(animals, &tiger)
	animals = append(animals, &human)

	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	weekDay := myRand.Intn(7)
	fmt.Println(weekDay)
	if weekDay > 0 && weekDay < 6 {
		for _, animal := range animals {
			animal.Live()
		}
	} else {
		//for _, animal := range animals {
		//	switch animal.(type) {
		//	case *Human:
		//		//将人转换为monster并令其捕食
		//		animal.(*Human).Hunt()
		//	case *Monster:
		//		animal.(*Monster).Run()
		//	default:
		//		animal.Death()
		//	}
		//}

		for _, animal := range animals {
			if human, ok := animal.(*Human); ok {
				human.Hunt()
			} else if tiger, ok := animal.(*BigTiger); ok {
				fmt.Println(tiger.Name)
			} else {
				animal.Death()
			}

		}
	}
}
