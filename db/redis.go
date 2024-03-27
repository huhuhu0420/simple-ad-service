package db

import (
	"context"
	"fmt"
	"time"

	"github.com/huhuhu0420/simple-ad-service/utils"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Cache interface {
	Set(context.Context, string, interface{}, time.Duration) *redis.StatusCmd
	Get(context.Context, string) *redis.StringCmd
	FlushAll(context.Context) *redis.StatusCmd
	Close() error
}

func InitRedis(config *utils.Config) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatal(err)
	}
	return rdb
}
