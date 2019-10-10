package main

import (
	"fmt"
	"sdk"
)

func main() {
	person := sdk.Person{
		Name:   "andre",
		Age:    12,
		Gender: "male",
	}
	person.GetInfo()
	fmt.Println(person.Age)
	//LoveManySkills(1.11, "Golang", "Rust", "Redis", "MySQL", "C")
	func(a, b int) {
		fmt.Println(int64(a + b))
	}(110100, 11019)

	var product = func(x, y float64) float64 {
		return x * y
	}
	fmt.Println(product(1101, 2))
}

// 不定长类型只能放在参数列表最后
func LoveManySkills(version float32, skills ...interface{}) {

	fmt.Println("current version is: ", version)
	for key, skill := range skills {
		fmt.Printf("你喜欢的第 %d 门技能是 %v\n", key+1, skill)
	}
}
