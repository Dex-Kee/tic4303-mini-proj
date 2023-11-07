package middleware

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

func BuildRedisClient(cfg *RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		DB:       cfg.DB,
		Password: cfg.Password,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Panicf("error when init redis client: %v", err)
	}

	return client
}
