package redis

import (
	"time"

	"github.com/go-redis/redis"

	"github.com/scheduler-service/src/config"
)

// Client : Redis client implementation
type Client struct {
	config config.RedisConfig
	client *redis.Client
}

// GetRedis function to return redis instance
func GetRedis(config config.RedisConfig) *Client {
	return &Client{config: config}
}

func (c *Client) Init() error {
	if c.client == nil {
		redisClient, err := c.genrateRedisClient()
		if err != nil {
			return err
		}
		c.client = redisClient
	}
	return nil
}

func (c *Client) genrateRedisClient() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.config.Host,
		ReadTimeout:  time.Duration(c.config.ReadTimeoutMillis) * time.Millisecond,
		WriteTimeout: time.Duration(c.config.WriteTimeoutMillis) * time.Millisecond,
	})

	return redisClient, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}

func (c *Client) Set(key string, value interface{}) error {
	return c.client.Set(key, value, -1).Err()
}

func (c *Client) Get(key string) (interface{}, error) {
	return c.client.Get(key).Result()
}

func (c *Client) Delete(key ...string) error {
	return c.client.Del(key...).Err()
}

func (c *Client) SetWithExpiry(key string, value interface{}, duration time.Duration) error {
	return c.client.Set(key, value, duration).Err()
}

func (c *Client) Publish(channel string, message interface{}) error {
	return c.client.Publish(channel, message).Err()
}
