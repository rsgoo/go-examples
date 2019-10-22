package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"os"
)

func HandleRedisError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleRedisError(err, "redis.Dia")

	defer conn.Close()

	Get(conn)

	Lrange(conn)

}

func Get(conn redis.Conn) {
	reply, err := conn.Do("get", "name")
	fmt.Printf("%T\n", reply)
	fmt.Printf("%v\n", reply)

	HandleRedisError(err, "conn.Do")
	replyStr, _ := redis.String(reply, err)
	fmt.Println(replyStr)
	fmt.Println(string(reply.([]byte)))
}

func Lrange(conn redis.Conn) {
	reply, err := redis.Values(conn.Do("LRANGE", "runoobkey", 0, -1))
	fmt.Printf("%T\n", reply)
	fmt.Printf("%v\n", reply)
	fmt.Println()

	HandleRedisError(err, "conn.Do")
	for _, value := range reply {
		fmt.Println(string(value.([]byte)))
		fmt.Println(redis.String(value, nil))
	}
}
