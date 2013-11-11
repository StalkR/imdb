package cache

import (
	"sync"
	"time"
)

// AppCache represents an in-app memory cache.
type AppCache struct {
	m     sync.Mutex // To protect items.
	items map[string]item
}

// item represents the internal type for items held in an AppCache.
type item struct {
	Value   []byte
	Expires time.Time
}

// NewAppCache creates a new AppCache.
func NewAppCache() AppCache {
	return AppCache{items: make(map[string]item)}
}

// Get gets a value given a key.
func (a AppCache) Get(key string) ([]byte, error) {
	a.m.Lock()
	defer a.m.Unlock()
	it, ok := a.items[key]
	if !ok {
		return nil, ErrCacheMiss
	}
	if !it.Expires.IsZero() && it.Expires.Before(time.Now()) {
		delete(a.items, key)
		return nil, ErrCacheMiss
	}
	return it.Value, nil
}

// Set sets a key to value with a default expiration of maximum one hour.
func (a AppCache) Set(key string, value []byte) {
	a.SetExpire(key, value, time.Hour)
}

// SetExpire sets a key to value with a given expiration.
// A zero value for expiration means it does not expire.
func (a AppCache) SetExpire(key string, value []byte, expiration time.Duration) {
	a.m.Lock()
	it := item{Value: value}
	if expiration > 0 {
		it.Expires = time.Now().Add(expiration)
	}
	a.items[key] = it
	a.m.Unlock()
	a.GC()
}

// GC garbage collects expired items from the cache.
func (a AppCache) GC() {
	a.m.Lock()
	defer a.m.Unlock()
	for key, item := range a.items {
		if item.Expires.Before(time.Now()) {
			delete(a.items, key)
		}
	}
}
