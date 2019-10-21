package main

import "fmt"

import (
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

func HandleSQLError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

//db 后对应数据表的字段
type Person struct {
	Name  string  `db:"name"`
	Age   int     `db:"age"`
	Money float64 `db:"rmb"`
}

func main() {
	//Create()
	//Delete()
	//Update()
	Select()
}

func Create() {
	db, err := sqlx.Open("mysql", "root:11019@tcp(localhost:3306)/mydb")
	HandleSQLError(err, "sqlx.Open")
	if err != nil {
		HandleSQLError(err, "sqlx.Open")
	}
	year := fmt.Sprintf("%v", time.Now().Year())

	result, err := db.Exec("insert into person(name, age,rmb,gender,birthday) values (?,?,?,?,?)", "rust", 12, 26000, true, year)
	if err != nil {
		HandleSQLError(err, "db.Exec")
	}

	fmt.Println(result.LastInsertId())
}

func Delete() {
	db, err := sqlx.Open("mysql", "root:11019@tcp(localhost:3306)/mydb")
	HandleSQLError(err, "sqlx.Open")
	if err != nil {
		HandleSQLError(err, "sqlx.Open")
	}

	result, err := db.Exec("DELETE from person where id > ?", 5)
	if err != nil {
		HandleSQLError(err, "db.Exec")
	}

	fmt.Println(result.RowsAffected())
}

func Update() {
	db, err := sqlx.Open("mysql", "root:11019@tcp(localhost:3306)/mydb")
	HandleSQLError(err, "sqlx.Open")
	if err != nil {
		HandleSQLError(err, "sqlx.Open")
	}

	result, err := db.Exec("update person set age = ? where id=4", 35)
	if err != nil {
		HandleSQLError(err, "db.Exec")
	}

	fmt.Println(result.RowsAffected())
}

func Select() {
	db, err := sqlx.Open("mysql", "root:11019@tcp(localhost:3306)/mydb")
	HandleSQLError(err, "sqlx.Open")
	if err != nil {
		HandleSQLError(err, "sqlx.Open")
	}

	people := []Person{}
	err = db.Select(&people, "SELECT name,age,rmb FROM  person")
	if err != nil {
		HandleSQLError(err, "db.Exec")
	}

	for _, item := range people {
		fmt.Println(item)
	}
}
