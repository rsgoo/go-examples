package demo_05

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (self *Student) init(name string, age, score int) {
	self.Name = name
	self.Age = age
	self.Score = score
}

func (self Student) get() Student {
	return self
}

func GetStructInfo()  {
	var nbaPlayer = Student{}
	nbaPlayer.init("Aldridge",33,23)

	cbaPlayer := nbaPlayer.get()
	fmt.Println(cbaPlayer)
}
