package library

import (
	"context"
	"fmt"
	"github.com/all-skeleton/gin-skeleton/config"
	"github.com/go-redis/redis/v8"
	"github.com/putyy/gokv"
	"time"
)

var redisObj *redis.Client
var redisKvCache = make(map[string]gokv.RedisIns)

func init() {
	redisObj = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password, // no password set
		DB:       config.Redis.Database, // use default DB
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := redisObj.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis err", err)
	}

	redisKvCache["user_info"] = createKv(map[string]string{"tail": "user_info"})
}

func Redis() *redis.Client {
	return redisObj
}

func BuildRdsKv(key string) gokv.RedisIns {
	return redisKvCache[key].(gokv.RedisIns)
}

func createKv(r map[string]string) gokv.RedisIns {
	kv := struct{ gokv.RedisKv }{gokv.RedisKv{gokv.Kv{}}}
	for i, v := range r {
		switch i {
		case "project":
			fallthrough
		case "pro":
			kv.RedisKv.Kv.Project = v
		case "tail":
			fallthrough
		case "t":
			kv.RedisKv.Kv.Tail = v
		case "Prefix":
			fallthrough
		case "pre":
			kv.RedisKv.Kv.Prefix = v
		case "separator":
			fallthrough
		case "sep":
			kv.RedisKv.Kv.Separator = v
		}
	}
	return kv
}
