package repository

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

type ChatRedis struct {
	ctx    *context.Context
	client *redis.Client
}

func NewChatRedis(ctx *context.Context) *ChatRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return &ChatRedis{
		ctx:    ctx,
		client: client,
	}
}

func (r *ChatRedis) Pong() (string, error) {
	return r.client.Ping(*r.ctx).Result()
}
