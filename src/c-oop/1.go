package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
	Hobby  []string
}

type Programmer struct {
	Person
	Device string
	Skills []string
}

func Info() {
	aaron := Person{
		Name:   "aaron",
		Age:    25,
		Gender: "male",
		Hobby:  []string{"basketball", "programing", "reading"},
	}

	linus := Programmer{
		Person: aaron,
		Device: "MBP",
		Skills: []string{
			"linux",
			"Nginx",
			"Golang",
			"MySQL",
			"Docker",
			"ElasticSearch",
			"Kafka",
		},
	}
	aaron.GetHobbies()
	linus.GetHobbies()
	linus.GetSkills()

}

func (p *Person) eat() {
	p.Name = "jerry"
	fmt.Println(p.Name, "可以吃饭")
}

func (p *Person) sport() {
	fmt.Println(p.Name, "在运动")
	p.Age--
}

func (p *Person) learn() {
	fmt.Println(p.Name, "在学习")
}

func (p *Person) age() {
	fmt.Println(p.Name, "的年龄是", p.Age)
}

func (p *Person) GetHobbies() {
	for _, value := range p.Hobby {
		fmt.Println(p.Name, "的兴趣爱好有:", value)
	}
}

func (coder *Programmer) GetHobbies() {
	for _, value := range coder.Hobby {
		fmt.Println(coder.Name, "的爱好有", value)
	}
}

func (coder *Programmer) GetSkills() {
	for _, value := range coder.Skills {
		fmt.Println(coder.Name, "的技能有:", value)
	}
}
