package chache

import (
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisIntegrationTest(t *testing.T) {
	redis := NewRedisCache(&RedisOption{Adrs: "localhost:6379"})
	err := redis.Set("1", "value1", 5*time.Minute)
	assert.NoError(t, err, "Expected no error on Set")

	val1, err := redis.Get("1")
	assert.NoError(t, err, "Expected no error on Get")
	assert.Equal(t, "value1", val1, "Expected retrieved value to match set value")

	err = redis.Delete("1")
	assert.NoError(t, err, "Expected no error on Delete")

	_, err = redis.Get("1")
	assert.Error(t, err, "Expected error for getting deleted key")
}

func TestRedisSet(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &Res{Client: db}

	key := "testKey"
	value := "testValue"
	ttl := time.Second * 10

	mock.ExpectSet(key, value, ttl).SetVal("OK")

	err := cache.Set(key, value, ttl)
	assert.NoError(t, err, "Expected no error on Set")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisGet(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &Res{Client: db}

	key := "testKey"
	expectedValue := "testValue"

	mock.ExpectGet(key).SetVal(expectedValue)

	value, err := cache.Get(key)
	assert.NoError(t, err, "Expected no error on Get")
	assert.Equal(t, expectedValue, value, "Expected retrieved value to match expected value")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisDelete(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &Res{Client: db}

	key := "testKey"

	mock.ExpectDel(key).SetVal(1)

	err := cache.Delete(key)
	assert.NoError(t, err, "Expected no error on Delete")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisClear(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &Res{Client: db}

	mock.ExpectFlushDB().SetVal("OK")

	err := cache.Clear()
	assert.NoError(t, err, "Expected no error on Clear")
	assert.NoError(t, mock.ExpectationsWereMet())
}
