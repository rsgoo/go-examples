package model

type Book struct {
	ID      int
	Title   string
	Author  string
	Price   float64
	Sales   int
	Stock   int //库存
	ImgPath string
	Ctime   string
	Mtime   string
}
