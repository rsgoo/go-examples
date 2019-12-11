package main

import (
	"net/http"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"z-web/model"
)

func main() {
	app := &model.App{ID: 1}
	//app.GetById()
	rows, _ := app.GetMany()

	for _, value := range rows {
		fmt.Println(*value)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello,world", r.URL.Path)
}
