package cache

import (
	"sync"
	"time"
)


type CacheItem struct {
	Value     interface{}
	CreatedAt time.Time
}

type Cache struct {
	data     map[string]CacheItem
	duration time.Duration
	mutex    sync.RWMutex
}

func NewCache(duration time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]CacheItem),
		duration: duration,
	}
	go c.cleanup()
	return c
}

func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = CacheItem{
		Value:     value,
		CreatedAt: time.Now(),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, found := c.data[key]
	if !found {
		return nil, false
	}

	if time.Since(item.CreatedAt) > c.duration {
		// delete(c.data, key)
		// return nil, false
		item.CreatedAt = time.Now()
		c.data[key] = item
	}
	
	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
}

func (c *Cache) cleanup() {
	for {
		select {
		case <-time.After(c.duration):
			c.mutex.Lock()
			for key, item := range c.data {
				if time.Since(item.CreatedAt) > c.duration {
					delete(c.data, key)
				}
			}
			c.mutex.Unlock()
		}
	}
}