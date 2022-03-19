package ch03

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

const (
	KeyValidWebsite = "app:valid:website:list"
)

func DoSomethingWithRedis(rdb *redis.Client, key string) bool {
	ctx := context.TODO()
	// func (c Client) SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd
	// Returns whether each member is a member of the set stored at key.
	if !rdb.SIsMember(ctx, KeyValidWebsite, key).Val() {
		return false
	}
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	if !strings.HasPrefix(val, "https://") {
		val = "https://" + val
	}
	if err := rdb.Set(ctx, "blog", val, 5*time.Second).Err(); err != nil {
		return false
	}
	return true
}
