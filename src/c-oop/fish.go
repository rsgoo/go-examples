package main

import "fmt"

type Fisher interface {
	Say() string
}

type BigFish struct {
}

func (b *BigFish) Say() (string) {
	return "a big fish"
}

type SmallFish struct {
}

func (s *SmallFish) Say() (string) {
	return "a small fish"
}

func main() {

	var big Fisher
	var small Fisher

	var fishList []Fisher
	big = &BigFish{}
	small = &SmallFish{}

	fishList = append(fishList, big, small)

	for key, value := range fishList {
		fmt.Println(key,":",value.Say())
	}
}
