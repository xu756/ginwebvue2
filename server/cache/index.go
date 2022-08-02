package cache

import (
	"context"
	"example.com/mod/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()
var RedisClient *redis.Client

func RedisInit() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.InitData.Redis.Addr,
		Password: config.InitData.Redis.Password,
		DB:       config.InitData.Redis.Db,
	})
	//测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}
	RedisClient = rdb
}

// Set 封装redis的set方法
func Set(key string, value interface{}, timeout int) {
	RedisInit()
	Del(key)
	err := RedisClient.Set(ctx, key, value, time.Duration(timeout)*time.Second).Err()
	if err != nil {
		fmt.Println("缓存错误")
		return
	}
}

// Get 封装redis的get方法
func Get(key string) interface{} {
	RedisInit()
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("缓存错误", err)
		return nil
	}
	return val
}

// Exists Exists 封装redis的exists方法
func Exists(key string) bool {
	RedisInit()
	res, _ := RedisClient.Exists(ctx, key).Result()
	if res == 1 {
		return true

	}
	return false
}

// Del 封装redis的del方法
func Del(key string) {
	RedisInit()
	err := RedisClient.Del(ctx, key).Err()
	if err != nil {
		fmt.Println("缓存错误")
		return
	}
}
