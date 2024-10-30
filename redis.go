package chache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Res struct {
	Client *redis.Client
}

type RedisOption struct {
	Adrs     string
	Username string
	Network  string
	DB       int
	Password string
}

func NewRedisCache(opt *RedisOption) *Res {
	RO := &redis.Options{
		Addr:     opt.Adrs,
		Username: opt.Username,
		Password: opt.Password,
		DB:       opt.DB,
	}
	return &Res{Client: redis.NewClient(RO)}
}

func (c *Res) Set(key string, value interface{}, ttl time.Duration) error {
	ctx := context.Background()
	return c.Client.Set(ctx, key, value, ttl).Err()
}

func (c *Res) Get(key string) (interface{},error) {
	ctx := context.Background()
	return c.Client.Get(ctx, key).Result()
}
func (c *Res) Delete(key string)error {
	ctx := context.Background()
	return c.Client.Del(ctx, key).Err()
}
func (c *Res) Clear() error {
	ctx := context.Background()
	return c.Client.FlushDB(ctx).Err()
}