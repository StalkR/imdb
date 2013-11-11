package cache

import (
	"sync"
	"time"
)

// An AppCache represents an in-app memory cache.
type AppCache struct {
	m     sync.Mutex // To protect items.
	items map[string]Item
}

// NewAppCache creates a new AppCache.
func NewAppCache() AppCache {
	return AppCache{items: make(map[string]Item)}
}

// Get gets a value given a key.
func (a AppCache) Get(key string) ([]byte, error) {
	a.m.Lock()
	defer a.m.Unlock()
	item, ok := a.items[key]
	if !ok {
		return nil, ErrCacheMiss
	}
	if !item.Expires.IsZero() && item.Expires.Before(time.Now()) {
		delete(a.items, key)
		return nil, ErrCacheMiss
	}
	return item.Value, nil
}

// Set sets a key to value with a default expiration of maximum one hour.
func (a AppCache) Set(key string, value []byte) {
	a.SetExpire(key, value, time.Hour)
}

// SetExpire sets a key to value with a given expiration.
// A zero value for expiration means it does not expire.
func (a AppCache) SetExpire(key string, value []byte, expiration time.Duration) {
	a.m.Lock()
	item := Item{Value: value}
	if expiration > 0 {
		item.Expires = time.Now().Add(expiration)
	}
	a.items[key] = item
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
