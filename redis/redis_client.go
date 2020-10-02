package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go_backend/config"
	"go_backend/log"
	"time"
)

var redisClient *redis.Client

const DefaultTimeOut = 100 * time.Millisecond

/**
 *
 * @Description  init the redis client
 * @Date 2:36 下午 2020/8/24
 **/
func init() {
	addr, err := config.GetValue("redis", "address")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisClient = rdb
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
func Set(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	err := redisClient.Set(ctx, key, value, 0).Err()
	return err

}
func SetNX(key string, value interface{}, ti time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	err := redisClient.SetNX(ctx, key, value, ti).Err()
	return err
}
func Expire(key string, ti time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	err := redisClient.Expire(ctx, key, ti).Err()
	return err
}
func HSet(key string, value ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	err := redisClient.HSet(ctx, key, value).Err()
	return err
}
func HGet(key string, fieldStr string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	val, err := redisClient.HGet(ctx, key, fieldStr).Result()
	return val, err
}

func HGetAll(key string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	val, err := redisClient.HGetAll(ctx, key).Result()
	return val, err
}
func Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	result, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Errorf("Get error:%s", err.Error())
		return "", err
	} else {
		return result, nil
	}
}
func TTL(key string) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	result, err := redisClient.TTL(ctx, key).Result()
	if err != nil {
		log.Errorf("Get error:%s", err.Error())
		return 0, err
	} else {
		return result, err
	}
}
func Exists(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeOut)
	defer cancel()
	result, err := redisClient.Exists(ctx, key).Result()
	if err != nil {
		log.Errorf("Get error:%s", err.Error())
		return 0, err
	} else {
		return result, err
	}
}
