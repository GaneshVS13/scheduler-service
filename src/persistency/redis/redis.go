package redis

import "time"

type Redis interface {
	Close() error
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key ...string) error
	SetWithExpiry(key string, value interface{}, duration time.Duration) error
	Publish(channel string, message interface{}) error
}
