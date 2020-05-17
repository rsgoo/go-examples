package main

import "fmt"

type I_Film interface {
	Make()
	OnShow()
}

type Film struct {
	Name       string
	LeadActors []string
}

//电影结构体实现IVideo接口的方法
func (f *Film) Make() {
	fmt.Println("film name is ", f.Name)
	fmt.Println("电影在制作")
}
func (f *Film) OnShow() {
	fmt.Println("电影在上映")
}

//电影子类
type EmotionFilm struct {
	Film
}

func (f *EmotionFilm) Make() {
	fmt.Println("情感电影在制作")
}

func main() {
	var myFilm I_Film
	var detailFilm = &Film{"WAR 2", []string{"actor1"}}
	myFilm = detailFilm
	myFilm.Make()
	myFilm.OnShow()
	fmt.Println(detailFilm.Name)

}
