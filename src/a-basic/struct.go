package main

import "fmt"

type Person struct {
	Name  string
	Age   int64
	Team  string
	Score int
}

func main() {
	var tom = Person{}
	tom.Name = "tom"
	tom.Age = 2111
	tom.Team = "tomTeam"
	tom.Score = 111
	fmt.Println(tom)

	fmt.Printf("%p\n", &tom)
	fmt.Printf("%p\n", &tom.Name)
	fmt.Printf("%p\n", &tom.Age)
	fmt.Printf("%p\n", &tom.Team)
	fmt.Printf("%p\n", &tom.Score)

}
