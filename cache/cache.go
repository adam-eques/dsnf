package main

import (
	"sync"
)

type Cache struct {
	mu      sync.Mutex
	records map[Key]*DNSrecord
}

type Key string

type DNSrecord struct {
	Name string
	IP   string
}

func NewCache() *Cache {
	return &Cache{
		records: make(map[Key]*DNSrecord),
	}
}

func (c *Cache) Insert(v DNSrecord) {
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.records[Key(v.Name)] = &v
}

func (c *Cache) Get(key Key) *DNSrecord {
	if key == "" {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if record, ok := c.records[key]; !ok {
		return record
	}

	return nil
}
