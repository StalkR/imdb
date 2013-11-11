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
	"appengine/datastore"
	"appengine/memcache"
)

// ErrCacheMiss is when Get is called and key is not found in any of the caches.
var ErrCacheMiss = errors.New("cache: miss")

// In-app memory cache.
var inAppCache = NewAppCache()

// An Item represents the internal type for items held in AppCache and datastore.
type Item struct {
	Value   []byte
	Expires time.Time
}

// Get gets a value given a key.
func Get(c appengine.Context, key string) ([]byte, error) {
	// First layer: in-app memory cache.
	value, err := inAppCache.Get(key)
	if err == nil {
		c.Infof("cache: in-app hit")
		return value, nil
	}
	// Second layer: memcache.
	item, err := memcache.Get(c, key)
	if err == nil {
		c.Infof("cache: memcache hit")
		return item.Value, nil
	}
	if err != memcache.ErrCacheMiss {
		c.Errorf("cache: memcache get: %v", err)
		return nil, err
	}
	// Third layer: datastore.
	k := datastore.NewKey(c, "Item", "Cache:"+key, 0, nil)
	e := Item{}
	err = datastore.Get(c, k, &e)
	if err == nil {
		if !e.Expires.IsZero() && e.Expires.Before(time.Now()) {
			c.Infof("cache: miss")
			return nil, ErrCacheMiss
		}
		c.Infof("cache: datastore hit")
		return e.Value, nil
	}
	if err != datastore.ErrNoSuchEntity {
		c.Errorf("cache: datastore get: %v", err)
		return nil, err
	}
	c.Infof("cache: miss")
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
		c.Errorf("cache: memcache set: %v", err)
		return err
	}
	// Third layer: datastore.
	// Datastore can save up to 1MB of []byte in an entity.
	if len(value) >= 1<<20 {
		return nil
	}
	k := datastore.NewKey(c, "Item", "Cache:"+key, 0, nil)
	e := Item{Value: value}
	if expiration > 0 {
		e.Expires = time.Now().Add(expiration)
	}
	if _, err := datastore.Put(c, k, &e); err != nil {
		c.Errorf("cache: datastore set: %v", err)
		return err
	}
	return nil
}
