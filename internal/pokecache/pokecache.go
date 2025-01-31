package pokecache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	records map[string]cacheRecord
	mutex   *sync.RWMutex
}

type cacheRecord struct {
	createdAt time.Time
	value     []byte
}

func CreateCache(duration time.Duration) *Cache {
	cache := Cache{
		mutex:   &sync.RWMutex{},
		records: map[string]cacheRecord{},
	}

	go cache.reapLoop(duration)

	return &cache
}

func (c Cache) Add(key string, value []byte) {
	defer log.Printf("New record has been added to cache with key: %s.\n", key)

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.records[key] = cacheRecord{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c Cache) Get(key string) (data []byte, wasFound bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if record, wasFound := c.records[key]; wasFound {
		log.Printf("Found cached data for key: %s\n", key)
		return record.value, wasFound
	} else {
		log.Printf("No cached data found for key: %s\n", key)
		return record.value, wasFound
	}
}

func (c Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		c.reap(time.Now(), duration)
	}
}

func (c Cache) reap(timestamp time.Time, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, record := range c.records {
		if timestamp.Add(duration).After(record.createdAt) {
			delete(c.records, key)
		}
	}
}
