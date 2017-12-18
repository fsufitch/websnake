package ui

import "time"

type cacheEntry struct {
	Data        []byte
	ContentType string
	Timestamp   time.Time
	Error       error
}

type cacheMap struct {
	Cache map[string]cacheEntry
	TTL   time.Duration
}

// newCacheMap creates a new proxy cache map
func newCacheMap(ttl time.Duration) *cacheMap {
	return &cacheMap{
		Cache: map[string]cacheEntry{},
		TTL:   ttl,
	}
}

// Set adds a successful value to the cache
func (p cacheMap) Set(key string, value []byte, contentType string) cacheEntry {
	p.Cache[key] = cacheEntry{
		Data:        value,
		ContentType: contentType,
		Timestamp:   time.Now(),
		Error:       nil,
	}
	return p.Cache[key]
}

// Set adds a failed value to the cache
func (p cacheMap) SetError(key string, err error) cacheEntry {
	p.Cache[key] = cacheEntry{
		Data:      []byte{},
		Timestamp: time.Now(),
		Error:     err,
	}
	return p.Cache[key]
}

// Delete removes a value from the cache
func (p cacheMap) Delete(key string) {
	delete(p.Cache, key)
}

// Get retrieves a cache entry
func (p cacheMap) Get(key string) (result cacheEntry, ok bool) {
	entry, ok := p.Cache[key]
	if ok {
		age := time.Now().Sub(entry.Timestamp)
		if age > p.TTL {
			p.Delete(key)
			ok = false
		} else {
			result = entry
		}
	}
	return
}
