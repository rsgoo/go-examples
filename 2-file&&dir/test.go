package main

import "fmt"

type Reader interface {
	Read() int
}

func NewReader(r Reader) (Reader) {
	return  r
}

type File struct {
}

func (f *File) Read() int {
	return 2
}

func main() {

	//var myReader Reader
	myReader := &File{}

	fmt.Println(NewReader(myReader).Read())
}
