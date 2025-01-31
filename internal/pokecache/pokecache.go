package pokecache

import (
	"sync"
	"time"
)

type cacheRecord struct {
	createdAt time.Time
	value     []byte
}

type cache struct {
	records map[string]cacheRecord
	mutex   *sync.RWMutex
}

func NewCache(duration time.Duration) *cache {
	cache := cache{
		mutex:   &sync.RWMutex{},
		records: map[string]cacheRecord{},
	}
	cache.invalidate(duration)

	return &cache
}

func (c cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.records[key] = cacheRecord{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c cache) Get(key string) (data []byte, wasFound bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	record, wasFound := c.records[key]

	if !wasFound {
		return nil, wasFound
	}

	return record.value, wasFound
}

func (c cache) invalidate(duration time.Duration) {
	ticker := time.NewTicker(duration)

	go func() {
		for {
			select {
			case tick := <-ticker.C:
				c.mutex.Lock()
				defer c.mutex.Unlock()

				for key, record := range c.records {
					expirationTime := record.createdAt.Add(duration)
					if tick.After(expirationTime) {
						delete(c.records, key)
					}
				}
			}
		}
	}()
}
