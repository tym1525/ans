package session

import (
	"github.com/garyburd/redigo/redis"
	"sync"
)

type RedisSession struct {
	pool   *redis.Pool
	rwlock sync.RWMutex
}

func NewRedisSession(id string) *RedisSession {
	return &RedisSession{}
}

func (r *RedisSession) Set(key string, value interface{}) (err error) {
	return
}

func (r *RedisSession) Get(key string) (value interface{}, err error) {

	return
}

func (r *RedisSession) Del(key string) (err error) {
	return
}

func (r *RedisSession) Save() (err error) {
	return
}
