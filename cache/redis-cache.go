package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/danish45007/go-rest/entity"
	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

// redis constructor function
func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

// create a redis client
func (cache *redisCache) CreateRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *entity.Post) {
	client := cache.CreateRedisClient()
	var ctx = context.Background()
	// serialization of post before insertion
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(ctx, key, json, cache.expires*time.Second)

}

func (cache *redisCache) Get(key string) *entity.Post {
	client := cache.CreateRedisClient()
	var ctx = context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}

	return &post
}
