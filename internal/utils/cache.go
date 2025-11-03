package utils

import (
	"context"
	"encoding/json"
	"payme/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	exp time.Duration
}

func NewRedisCache(exp time.Duration) (*RedisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.CFG.REDIS_ADDR,
		Password: config.CFG.REDIS_PASS,
		DB: 0, // TODO: what if new connection?
	})

	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		rdb.Close();
		return nil, err
	}

	return &RedisCache{
		client: rdb,
		exp: exp,
	}, nil
}

func (r *RedisCache) CloseDb() {
	r.client.Close();
}

func (r *RedisCache)Get(key string) (any, error) {
	data, err := r.client.Get(context.TODO(), key).Result()
	if err != nil {
		return nil, err
	}

	var body any
	if err := json.Unmarshal([]byte(data), &body); err != nil {
		return nil, err
	}

	return body, nil
}

func (r *RedisCache)Set(key string, value any) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = r.client.SetEx(context.TODO(), key, body, r.exp).Result()
	return err
}
