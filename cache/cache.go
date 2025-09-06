package cache

import (
	"crypto/sha256"
	"errors"
	"time"
)

var (
	CacheExpiration time.Duration = 12 * time.Hour // Max time that a request can exist in the cache
	MaxCacheMemory  int           = 500            // Max request can exist in the cache
)

// Find and delete string s in string slice
func FindAndDel(arr []string, s string) []string {
	index := 0
	for i, v := range arr {
		if v == s {
			index = i
			break
		}
	}
	return append(arr[:index], arr[index+1:]...)
}

/*
Create a cache that store:

- Key: API request URL

- Value: RequestResponse
*/
func MakeWikiCache() WikiCache {
	res := WikiCache{}
	res.Memory = map[string][]byte{}
	res.HashedKeyQueue = make([]string, 0, MaxCacheMemory)
	res.CreatedTime = map[string]time.Time{}
	return res
}

// Cache to store Wikipedia request result
type WikiCache struct {
	Memory         map[string][]byte    // Map store request result
	HashedKeyQueue []string             // Key queue. Delete the first item if reach max cache
	CreatedTime    map[string]time.Time // Map store created time
}

// Hash a string into SHA256
func HashCacheKey(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	return string(hasher.Sum(nil))
}

// Get WikiCache current number of cache
func (cache WikiCache) GetLen() int {
	return len(cache.HashedKeyQueue)
}

// Add cache into the WikiCache
func (cache *WikiCache) Add(s string, res []byte) {
	if len(cache.Memory) >= MaxCacheMemory {
		cache.Pop()
	}
	key := HashCacheKey(s)
	if cache.Memory == nil {
		cache.Memory = map[string][]byte{}
		cache.CreatedTime = map[string]time.Time{}
		cache.HashedKeyQueue = make([]string, 0, MaxCacheMemory)
	}
	if _, ok := cache.Memory[key]; !ok {
		cache.Memory[key] = res
		cache.CreatedTime[key] = time.Now()
		cache.HashedKeyQueue = append(cache.HashedKeyQueue, key)
	}
}

// Get response from the Cache
func (cache *WikiCache) Get(s string) ([]byte, error) {
	key := HashCacheKey(s)
	if value, ok := cache.Memory[key]; ok {
		if time.Since(cache.CreatedTime[key]) <= CacheExpiration {
			cache.HashedKeyQueue = FindAndDel(cache.HashedKeyQueue, key)
			cache.HashedKeyQueue = append(cache.HashedKeyQueue, key)
			return value, nil
		} else {
			cache.HashedKeyQueue = FindAndDel(cache.HashedKeyQueue, key)
			delete(cache.Memory, key)
			return []byte{}, errors.New("the data is outdated")
		}
	}
	return []byte{}, errors.New("cache key not exist")
}

// Delete the first key in the Cache
func (cache *WikiCache) Pop() {
	if len(cache.HashedKeyQueue) == 0 {
		return
	}
	delete(cache.Memory, cache.HashedKeyQueue[0])
	cache.HashedKeyQueue = cache.HashedKeyQueue[1:]
}

// Clear the whole Cache
func (cache *WikiCache) Clear() {
	*cache = WikiCache{}
	// This line to avoid declare but not used error
	_ = cache
}
