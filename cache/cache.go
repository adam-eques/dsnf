package cache

import (
	"fmt"
	"sync"
)

var fileCache = "file.cache"

type Cache struct {
	Lock  sync.Mutex
	Rlock sync.RWMutex
	cache map[string]string
}

func NewCache() (*Cache, error) {
	cache := &Cache{
		cache: make(map[string]string),
	}

	return cache, nil
}

func (c *Cache) Set(key, value string) error {
	c.Lock.Lock()
	defer c.Rlock.Unlock()

	if key == "" && value == "" {
		return fmt.Errorf("Invalid key/value data")
	}

	c.cache[key] = value

	return nil
}

func (c *Cache) Get(key string) (string, error) {
	var err error

	if key == "" {
		return "", fmt.Errorf("Empty key found")
	}

	c.Rlock.RLock()
	v, ok := c.cache[key]
	if !ok {
		return "", fmt.Errorf("Key does not exist: %v", err)
	}

	c.Rlock.RUnlock()
	return v, nil
}
