package sdk

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p *Person) GetInfo() {
	p.Age++
	fmt.Printf("name:%v, age:%d, gender:%v\n",p.Name,p.Age,p.Gender)
}
