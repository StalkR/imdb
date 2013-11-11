/*
Package cache implements an abstraction of caching for slice of bytes.

There are various techniques for caching on AppEngine, from fastest to slowest:
 - in-app memory cache
 - memcache (currently implemented)
 - datastore
*/
package cache

import (
	"time"

	"appengine"
	"appengine/memcache"
)

// Get gets a value given a key.
func Get(c appengine.Context, key string) ([]byte, error) {
	data, err := memcache.Get(c, key)
	if err != nil {
		if err != memcache.ErrCacheMiss {
			c.Errorf("cache: get: %v", err)
		}
		return nil, err
	}
	return data.Value, nil
}

// Set sets a key to value with a default expiration of maximum one hour.
func Set(c appengine.Context, key string, value []byte) error {
	return SetExpire(c, key, value, time.Hour)
}

// SetExpire sets a key to value with a given expiration.
func SetExpire(c appengine.Context, key string, value []byte, expiration time.Duration) error {
	item := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
	if err := memcache.Set(c, item); err != nil {
		c.Errorf("cache: set: %v", err)
		return err
	}
	return nil
}
