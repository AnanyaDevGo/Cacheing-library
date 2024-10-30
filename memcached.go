package chache

import (
	"encoding/json"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type memcached struct {
	client *memcache.Client
}

func NewMemcachedCache(server string) *memcached {
	return &memcached{client: memcache.New(server)}
}

func (c *memcached) Set(key string, value interface{}, ttl time.Duration) error {

	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(&memcache.Item{
		Key:        key,
		Value:      valueBytes,
		Expiration: int32(ttl.Seconds()),
	})
}

func (c *memcached) Get(key string) (string, error) {
	item, err := c.client.Get(key)
	if err != nil {
		return "", err
	}

	var value string
	if err := json.Unmarshal(item.Value, &value); err != nil {
		return "", err
	}
	return value, nil
}

func (c *memcached) Delete(key string) error {
	return c.client.Delete(key)
}

func (c *memcached) Clear() error {
	return c.client.DeleteAll()
}
