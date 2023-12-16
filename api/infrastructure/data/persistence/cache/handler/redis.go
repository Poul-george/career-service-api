package handler

import (
	"github.com/Poul-george/go-api/api/config"
	"github.com/go-redis/redis/v8"
)

type RedisHandler struct {
	client *redis.Client
}

func NewCacheHandler(c config.Redis) *RedisHandler {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Port,
		Password: "",          // no password set
		DB:       c.SessionDB, // use default DB
	})

	return &RedisHandler{
		client: rdb,
	}
}
