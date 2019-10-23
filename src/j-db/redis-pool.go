package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"strconv"
	"fmt"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:   20, //最大闲置连接数
		MaxActive: 0,  //最大活动连接数,0表示无限制
		//闲置连接的超时时间
		IdleTimeout: time.Second * 60,

		//定义拨号器获取连接的函数
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "127.0.0.1:6379")
			return conn, err
		},
	}

	defer pool.Close()

	for i := 0; i < 100; i++ {
		go func(pool *redis.Pool, i int) {
			//去连接池中去一个连接
			conn := pool.Get()
			defer conn.Close()
			reply, _ := conn.Do("set", "conn"+strconv.Itoa(i), i*i)
			str, _ := redis.String(reply, nil)
			fmt.Println(conn)
			fmt.Println(i)
			fmt.Println(str)
			fmt.Println("-----------------")
		}(pool, i)
	}

	time.Sleep(time.Second * 1)

}
