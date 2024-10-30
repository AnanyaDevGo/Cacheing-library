package chache

import "time"

type CacheBackend interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string)(interface{}, error)
	Delete(key string) error
	clear() error
}