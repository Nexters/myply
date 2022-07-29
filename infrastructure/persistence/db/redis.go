package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/go-redis/redis/v9"
)

type RedisInstance struct {
	Client  *redis.Client
	timeout time.Duration
}

func NewRedisDB(config *configs.Config) (*RedisInstance, error) {
	clientOptions, err := redis.ParseURL(config.RedisURI)

	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(clientOptions)

	return &RedisInstance{Client: rdb, timeout: config.RedisTimeout}, nil
}

func (rdb *RedisInstance) SetNXJson(key string, data interface{}, ttl time.Duration) error {
	b, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		return jsonErr
	}

	ctx, cancel := rdb.createCtx()
	defer cancel()

	_, err := rdb.Client.SetNX(ctx, key, b, ttl).Result()

	return err
}

func (rdb *RedisInstance) GetJson(key string, model interface{}) error {
	ctx, cancel := rdb.createCtx()
	defer cancel()

	b, err := rdb.Client.Get(ctx, key).Bytes()

	if err != nil {
		return err
	}

	json.Unmarshal(b, model)

	return nil
}

func (rdb *RedisInstance) ZIncrByOne(key string, member string) (float64, error) {
	ctx, cancel := rdb.createCtx()
	defer cancel()

	return rdb.Client.ZIncrBy(ctx, key, 1, member).Result()
}

func (rdb *RedisInstance) ZRevRange(key string, start int64, end int64) ([]string, error) {
	ctx, cancel := rdb.createCtx()
	defer cancel()

	return rdb.Client.ZRevRange(ctx, key, start, end).Result()
}

func (rdb *RedisInstance) createCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), rdb.timeout)
}
