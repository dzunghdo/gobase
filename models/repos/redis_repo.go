package repos

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gobase/db"
	"time"
)

type RedisRepository struct {
	redisClient redis.Client
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{redisClient: *db.GetRedisClient()}
}

func (repo *RedisRepository) Set(ctx context.Context, key string, value string) error {
	return repo.redisClient.Set(ctx, key, value, 0).Err()
}

func (repo *RedisRepository) SetWithDuration(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return repo.redisClient.Set(ctx, key, value, duration).Err()
}

func (repo *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	return repo.redisClient.Get(ctx, key).Result()
}
