package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	Rdb = NewRedisClient()
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	return rdb
}