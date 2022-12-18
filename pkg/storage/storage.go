package storage

import (
	"context"
	"net"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/po1yb1ank/ccounter/config"
)

type IStorage interface {
	Increment(ctx context.Context, key string) (int64, error)
	Decrement(ctx context.Context, key string) (int64, error)
	Set(ctx context.Context, key string, value int64) error
	Current(ctx context.Context, key string) (int64, error)
}

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(config config.Redis) IStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	return &RedisStorage{
		client: client,
	}
}

func (r *RedisStorage) Increment(ctx context.Context, key string) (int64, error) {
	if r.client.Exists(ctx, key).Val() == 0 {
		return 0, ErrorKeyNotFound
	}

	result, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *RedisStorage) Decrement(ctx context.Context, key string) (int64, error) {
	if r.client.Exists(ctx, key).Val() == 0 {
		return 0, ErrorKeyNotFound
	}

	result, err := r.client.Decr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *RedisStorage) Set(ctx context.Context, key string, value int64) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisStorage) Current(ctx context.Context, key string) (int64, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, ErrorKeyNotFound
	}

	val, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return int64(val), nil
}
