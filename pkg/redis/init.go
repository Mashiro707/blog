package redis

import (
	"blog/config"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
}

var RedisClient *redis.Client

func (r *Redis) Setup() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.RedisConfig.Addr,
		Password:     config.RedisConfig.Password,
		DB:           config.RedisConfig.DB,
		MinIdleConns: config.RedisConfig.MinIdleConn,
		DialTimeout:  config.RedisConfig.DialTimeout * time.Second,
		ReadTimeout:  config.RedisConfig.ReadTimeout * time.Second,
		WriteTimeout: config.RedisConfig.WriteTimeout * time.Second,
		PoolSize:     config.RedisConfig.PoolSize,
		PoolTimeout:  config.RedisConfig.PoolTimeout * time.Second,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
