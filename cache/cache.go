package cache

import (
	"crypto/sha256"
	"errors"
	"time"

	"github.com/trietmn/go-wiki/models"
)

var (
	CacheExpiration time.Duration = 12 * time.Hour
	MaxCacheMemory  int           = 500
)

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

func MakeWikiCache() WikiCache {
	res := WikiCache{}
	res.Memory = map[string]models.RequestResult{}
	res.HashedKeyQueue = make([]string, 0, MaxCacheMemory)
	res.CreatedTime = map[string]time.Time{}
	return res
}

type WikiCache struct {
	Memory         map[string]models.RequestResult
	HashedKeyQueue []string
	CreatedTime    map[string]time.Time
}

func HashCacheKey(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	return string(hasher.Sum(nil))
}

func (cache WikiCache) GetLen() int {
	return len(cache.HashedKeyQueue)
}

func (cache *WikiCache) Add(s string, res models.RequestResult) {
	if len(cache.Memory) >= MaxCacheMemory {
		cache.Pop()
	}
	key := HashCacheKey(s)
	if cache.Memory == nil {
		cache.Memory = map[string]models.RequestResult{}
		cache.CreatedTime = map[string]time.Time{}
		cache.HashedKeyQueue = make([]string, 0, MaxCacheMemory)
	}
	if _, ok := cache.Memory[key]; !ok {
		cache.Memory[key] = res
		cache.CreatedTime[key] = time.Now()
		cache.HashedKeyQueue = append(cache.HashedKeyQueue, key)
	}
}

func (cache *WikiCache) Get(s string) (models.RequestResult, error) {
	key := HashCacheKey(s)
	if value, ok := cache.Memory[key]; ok {
		if time.Since(cache.CreatedTime[key]) <= CacheExpiration {
			cache.HashedKeyQueue = FindAndDel(cache.HashedKeyQueue, key)
			cache.HashedKeyQueue = append(cache.HashedKeyQueue, key)
			return value, nil
		} else {
			cache.HashedKeyQueue = FindAndDel(cache.HashedKeyQueue, key)
			delete(cache.Memory, key)
			return models.RequestResult{}, errors.New("the data is outdated")
		}
	}
	return models.RequestResult{}, errors.New("cache key not exist")
}

func (cache *WikiCache) Pop() {
	if len(cache.HashedKeyQueue) == 0 {
		return
	}
	delete(cache.Memory, cache.HashedKeyQueue[0])
	cache.HashedKeyQueue = cache.HashedKeyQueue[1:]
}

func (cache *WikiCache) Clear() {
	*cache = WikiCache{}
	// This line to avoid declare but not used error
	_ = cache
}
