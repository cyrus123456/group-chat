package model

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

var RedisClientPool *redis.Client

func InitRedisclientpool() (err error) {
	RedisClientPool = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RedisClientPool.Ping(ctx).Result()
	return
}
