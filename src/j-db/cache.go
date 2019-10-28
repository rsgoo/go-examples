package main

import (
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

//db 后对应数据表的字段
type Human struct {
	Name  string  `db:"name"`
	Age   int     `db:"age"`
	Money float64 `db:"rmb"`
}

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	var cmd string
	for {
		fmt.Println("input a redis command...")
		fmt.Scan(&cmd)

		switch strings.ToLower(cmd) {
		case "getall":
			GetAllPeople()
		case "get":
			GetOne()
		case "exit":
			goto GAME_OVER
		default:
			fmt.Println("请输入有效命令")
		}
	}
GAME_OVER:
	fmt.Println("Game over")

}

func GetAllPeople() {
	fmt.Println("获取getallpeople")
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleError(err, "redis.Dial")
	//先尝试取缓存数据
	defer conn.Close()

	ok, data := GetPeopleFromRedis()
	if ok {
		fmt.Println("从缓存获取到的值是 ")
		for _, value := range data {
			fmt.Println(value)
		}
	} else {
		people := GetPeopleFromMySQL()
		CacheToRedis(&people)
	}

}

func GetOne() {
	fmt.Println("获取其中之一")
}

func CacheToRedis(human *[]Human) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleError(err, "redis.Dia")

	defer conn.Close()
	for _, h := range *human {
		humanStr := fmt.Sprint(h)
		_, err := conn.Do("lpush", "people", humanStr)
		if err != nil {
			fmt.Println("缓存失败 err = ", err)
		}
	}

	_, err1 := conn.Do("expire", "people", 300)
	if err1 != nil {
		HandleError(err1, "conn.Do expire")
	}
	fmt.Println("缓存成功")
}

//从缓存中取数据
func GetPeopleFromRedis() (ok bool, data []string) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleError(err, "redis.Dial")
	reply, err := conn.Do("lrange", "people", "0", "-1")
	if err != nil {
		HandleError(err, "lrange error")
	}

	//peopleStr := string(reply.([]byte))
	peopleStr, err := redis.Strings(reply, err)
	if err != nil {
		HandleError(err, "redis string")
	}

	if len(peopleStr) > 0 {
		return true, peopleStr
	}
	return false, nil

}

func GetPeopleFromMySQL() (people []Human) {
	db, err := sqlx.Open("mysql", "root:11019@tcp(localhost:3306)/mydb")
	defer db.Close()

	if err != nil {
		HandleError(err, "sqlx.Open")
		os.Exit(1)
	}

	//var people []Human
	err = db.Select(&people, "SELECT name,age,rmb FROM  person where id >= ?", 1)
	if err != nil {
		HandleError(err, "db.Select")
	}

	return  people
}
