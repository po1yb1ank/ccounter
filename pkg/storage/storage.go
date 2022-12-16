package storage

import (
	"context"
	"errors"
	"net"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/po1yb1ank/ccounter/config"
)

type IStorage interface {
	Increment(ctx context.Context, key string) (int64, error)
	Decrement(ctx context.Context, key string) (int64, error)
	Reset(ctx context.Context, key string) error
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
	result, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *RedisStorage) Decrement(ctx context.Context, key string) (int64, error) {
	result, err := r.client.Decr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *RedisStorage) Reset(ctx context.Context, key string) error {
	return r.client.Set(ctx, key, 0, 0).Err()
}

func (r *RedisStorage) Current(ctx context.Context, key string) (int64, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if errors.Is(err, redis.Nil) {
		_, err := r.client.Set(ctx, key, 0, 0).Result()
		if err != nil {
			return 0, err
		}

		return 0, nil
	}

	val, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return int64(val), nil
}
