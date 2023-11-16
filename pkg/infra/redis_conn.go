package infra

import (
	"go-gacha-system/pkg/config"

	"github.com/go-redis/redis/v8"
)

type RedisConnector struct {
	Client *redis.Client
}

func NewRedisConnector() *RedisConnector {
	conf := config.LoadRedisConfig()
	c := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: "",
		DB:       0,
	})
	return &RedisConnector{
		Client: c,
	}
}
