package chache

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type inMemory struct {
	maxSize int
	items   map[string]*list.Element
	order   *list.List
	mu      sync.RWMutex
}

type cacheNode struct {
	key       string
	value     interface{}
	expiresAt time.Time
}

func NewInMemoryCache(maxSize int) *inMemory {
	cache := &inMemory{
		maxSize: maxSize,
		items:   make(map[string]*list.Element),
		order:   list.New(),
	}

	go cache.expiryWorker()
	return cache
}

func (c *inMemory) Set(key string, value interface{}, ttl time.Duration) error {

	c.mu.RLock()
	elem, exists := c.items[key]
	c.mu.RUnlock()

	if exists {
		c.mu.Lock()
		node := elem.Value.(*cacheNode)
		node.value = value
		if ttl > 0 {
			node.expiresAt = time.Now().Add(ttl)
		}
		c.order.MoveToFront(elem)
		c.mu.Unlock()
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.order.Len() >= c.maxSize {
		c.evict()
	}

	expiresAt := time.Time{}
	if ttl > 0 {
		expiresAt = time.Now().Add(ttl)
	}

	newNode := &cacheNode{key: key, value: value, expiresAt: expiresAt}
	elem = c.order.PushFront(newNode)
	c.items[key] = elem
	return nil
}

func (c *inMemory) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	elem, exists := c.items[key]
	if !exists {
		return nil, errors.New("key not found")
	}

	node := elem.Value.(*cacheNode)
	if !node.expiresAt.IsZero() && time.Now().After(node.expiresAt) {
		return nil, errors.New("key expired")
	}

	c.order.MoveToFront(elem)
	return node.value, nil
}

func (c *inMemory) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, exists := c.items[key]
	if !exists {
		return errors.New("key not found")
	}

	c.order.Remove(elem)
	delete(c.items, key)
	return nil
}

func (c *inMemory) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*list.Element)
	c.order.Init()
}

func (c *inMemory) evict() {
	last := c.order.Back()
	if last != nil {
		node := last.Value.(*cacheNode)
		delete(c.items, node.key)
		c.order.Remove(last)
	}
}
func (c *inMemory) expiryWorker() {
	for {
		time.Sleep(time.Second) 
		c.mu.Lock()
		for {
			elem := c.order.Back()
			if elem == nil {
				break
			}
			node := elem.Value.(*cacheNode)
			if !node.expiresAt.IsZero() && time.Now().After(node.expiresAt) {
				c.order.Remove(elem)
				delete(c.items, node.key)
			} else {
				break 
			}
		}
		c.mu.Unlock()
	}
}
