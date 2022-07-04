package redis

import (
	"blog/config"
	"context"
	"github.com/go-redis/redis/extra/redisotel/v8"
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
		DialTimeout:  config.RedisConfig.DialTimeout,
		ReadTimeout:  config.RedisConfig.ReadTimeout,
		WriteTimeout: config.RedisConfig.WriteTimeout,
		PoolSize:     config.RedisConfig.PoolSize,
		PoolTimeout:  config.RedisConfig.PoolTimeout,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	if config.RedisConfig.EnableTrace {
		RedisClient.AddHook(redisotel.NewTracingHook())
	}
}
