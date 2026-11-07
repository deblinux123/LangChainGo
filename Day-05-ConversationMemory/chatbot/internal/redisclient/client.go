package redisclient

import (
	"context"
	"fmt"
	"time"

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

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

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

func (c *Client) SaveHistory(
	ctx context.Context,
	sessionID string,
	data []byte,
) error {
	key := "chat:session:" + sessionID

	if err := c.rdb.Set(
		ctx,
		key,
		data,
		0,
	).Err(); err != nil {
		return fmt.Errorf("Save history: %w", err)
	}
	return nil
}

func (c *Client) GetHistory(
	ctx context.Context,
	sessionID string,
) ([]byte, error) {
	key := "chat:session:" + sessionID

	data, err := c.rdb.Get(
		ctx,
		key,
	).Bytes()

	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, fmt.Errorf("get history: %w", err)
	}

	return data, nil
}

func (c *Client) DeleteHistory(
	ctx context.Context,
	sessionID string,
) error {
	key := "chat:session:" + sessionID

	if err := c.rdb.Del(
		ctx,
		key,
	).Err(); err != nil {
		return fmt.Errorf("delete history: %w", err)
	}

	return nil
}

func (c *Client) Close() error {
	return c.rdb.Close()
}
