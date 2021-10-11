package scheduler

import (
	"time"

	"github.com/scheduler-service/src/persistency/redis"
)

var (
	lockDurationInSec = 60
)

// locker implementation with Redis
type locker struct {
	client redis.Client
}

func (s *locker) Lock(key string) (bool, error) {
	err := s.client.SetWithExpiry(key, time.Now().String(), time.Second*time.Duration(lockDurationInSec))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *locker) Unlock(key string) error {
	return s.client.Delete(key)
}
