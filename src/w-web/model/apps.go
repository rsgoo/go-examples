package model

import (
	"w-web/utils"
	"fmt"
)

type App struct {
	ID       int    `db:"id"`
	App_name string `db:"app_name"`
	Url      string `db:"url"`
	Country  string `db:"country"`
	Name     string `db:"name"`
}

//预处理方式
func (a *App) Create() error {
	sqlStr := "insert into App(app_name, url, country, name) values(?,?,?,?)"

	//预处理
	stmt, err := utils.DB.Prepare(sqlStr)

	if err != nil {
		fmt.Println("DB.Prepare,", err.Error())
		return err
	}

	_, err = stmt.Exec("github", "www.github.com", "us", "github")
	if err != nil {
		fmt.Println("stmt.Exec,", err.Error())
		return err
	}
	return nil
}

//直接执行模式
func (a *App) Add() error {
	sqlStr := "insert into App(app_name, url, country, name) values(?,?,?,?)"

	//预处理
	_, err := utils.DB.Exec(sqlStr, "bing", "www.bing.com", "cn", "bing")

	if err != nil {
		fmt.Println("DB.Exec,", err.Error())
		return err
	}

	return nil
}

//查询单条记录
func (a *App) GetOne() error {
	fmt.Println(a)
	sqlStr := "select * from App where id = ?"
	row := utils.DB.QueryRow(sqlStr, a.ID)
	var id int
	var app_name string
	var url string
	var country string
	var name string
	err := row.Scan(&id, &app_name, &url, &country, &name)
	if err != nil {
		return err
	}

	a.ID = id
	a.App_name = app_name
	a.Url = url
	a.Country = country
	a.Name = name
	//App := App{
	//	ID:       id,
	//	App_name: app_name,
	//	Url:      url,
	//	Country:  country,
	//	Name:     name,
	//}
	return nil
}

//查询多条记录
func (a *App) GetMany() ([]*App, error) {
	sqlStr := "select * from apps"
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var apps []*App
	for rows.Next() {
		var id int
		var app_name string
		var url string
		var country string
		var name string

		err = rows.Scan(&id, &app_name, &url, &country, &name)
		if err != nil {
			return nil, err
		}

		record := &App{
			ID:       id,
			App_name: app_name,
			Url:      url,
			Country:  country,
			Name:     name,
		}

		apps = append(apps, record)
	}
	return apps, nil
}
