package database

import (
	"context"

	"canteen-backend/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Redis struct {
	client *redis.Client
}

var RedisClient *Redis

func NewRedisClient() *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		println("[redis] Connected to database")
	}

	RedisClient = &Redis{client}
	return RedisClient
}
