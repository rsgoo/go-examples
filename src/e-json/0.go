package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Name    string
	Age     int
	Team    string
	Country string
}

func main() {

	var White = Player{
		Name:    "White",
		Age:     23,
		Team:    "Spurs",
		Country: "USA",
	}

	var Mills = Player{
		Name:    "Mills",
		Age:     30,
		Team:    "Spurs",
		Country: "Australia",
	}

	Players := make([]Player, 0)

	Players = append(Players, White)
	Players = append(Players, Mills)

	//结构体切片转json字符串
	playersBytes, err := json.Marshal(Players)
	if err != nil {
		fmt.Println("err is ", err)
	}
	playJsonStr := string(playersBytes)
	fmt.Println(playJsonStr)

	//json字符串转结构体切片
	decodePlayers := make([]Player, 0)
	playStrBytes := []byte(playJsonStr)

	err1 := json.Unmarshal(playStrBytes, &decodePlayers)
	if err1 != nil {
		fmt.Println("err1 is ", err1)
	}

	fmt.Println(decodePlayers)

}
