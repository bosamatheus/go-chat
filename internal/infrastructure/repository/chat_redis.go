package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type ChatRedis struct {
	ctx    context.Context
	client *redis.Client
}

func NewChatRedis(ctx context.Context, client *redis.Client) *ChatRedis {
	return &ChatRedis{
		ctx:    ctx,
		client: client,
	}
}

func (r *ChatRedis) KeyNotExists(key string) bool {
	_, err := r.client.Get(r.ctx, key).Result()
	return err == redis.Nil
}

func (r *ChatRedis) GetPreviousValues(key string) ([]string, error) {
	return r.client.LRange(r.ctx, key, 0, -1).Result()
}

func (r *ChatRedis) SaveValue(key string, val []byte) error {
	return r.client.RPush(r.ctx, key, val).Err()
}
