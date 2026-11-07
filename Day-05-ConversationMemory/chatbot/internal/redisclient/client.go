package redisclient

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

type Client struct {
	rdb *redis.Client
}

func NewClient(cfg Config) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		_ = rdb.Close()

		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	return &Client{
		rdb: rdb,
	}, nil
}

func (c *Client) GetClient() *redis.Client {
	return c.rdb
}

func (c *Client) Close() error {
	return c.rdb.Close()
}
