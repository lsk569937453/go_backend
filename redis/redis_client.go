package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go_backend/config"
	"go_backend/log"
)

var ctx = context.Background()

var redisClient *redis.Client

/**
 *
 * @Description  init the redis client
 * @Date 2:36 下午 2020/8/24
 **/
func init() {
	addr := config.GetValue("redis", "address")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisClient = rdb
}
func Set(key string, value string) error {
	err := redisClient.Set(ctx, key, value, 0).Err()
	return err

}
func Get(key string) string {
	result, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Error("Get error:", err.Error())
		return ""
	} else {
		return result
	}

}
