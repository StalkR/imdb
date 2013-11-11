/*
Package cache implements caching of bytes on AppEngine in three layers.

It does three layers of caching from fastest to slowest:
 - in-app memory cache
 - memcache
 - datastore (not yet implemented)

If you only need an in-app memory cache, its implementation is exported so
you can use it directly.
*/
package cache

import (
	"errors"
	"time"

	"appengine"
	"appengine/memcache"
)

// ErrCacheMiss is when Get is called and key is not found in any of the caches.
var ErrCacheMiss = errors.New("cache: miss")

// In-app memory cache.
var inAppCache = NewAppCache()

// Get gets a value given a key.
func Get(c appengine.Context, key string) ([]byte, error) {
	// First layer: in-app memory cache.
	value, err := inAppCache.Get(key)
	if err == nil {
		return value, nil
	}
	// Second layer: memcache.
	item, err := memcache.Get(c, key)
	if err == nil {
		return item.Value, nil
	}
	if err != memcache.ErrCacheMiss {
		c.Errorf("cache: get: %v", err)
		return nil, err
	}
	// TODO(StalkR): Third layer: datastore.
	return nil, ErrCacheMiss
}

// Set sets a key to value with a default expiration of maximum one hour.
func Set(c appengine.Context, key string, value []byte) error {
	return SetExpire(c, key, value, time.Hour)
}

// SetExpire sets a key to value with a given expiration.
func SetExpire(c appengine.Context, key string, value []byte, expiration time.Duration) error {
	// First layer: in-app memory cache.
	inAppCache.SetExpire(key, value, expiration)
	// Second layer: memcache.
	item := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
	if err := memcache.Set(c, item); err != nil {
		c.Errorf("cache: set: %v", err)
		return err
	}
	// TODO(StalkR): Third layer: datastore.
	return nil
}
