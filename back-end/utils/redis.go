package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func InitRedis() {
	url := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("Invalid Redis URL: %v", err)
	}
	Redis = redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = Redis.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}
}

func GetCached(ctx context.Context, key string) string {
	val, err := Redis.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

func SetCached(ctx context.Context, key, value string, ttl time.Duration) {
	Redis.Set(ctx, key, value, ttl)
}
