package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Init(config *Config) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         config.Addr(),
		Username:     config.User,
		Password:     config.Password,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout * time.Second,
		ReadTimeout:  config.ReadTimeout * time.Second,
		WriteTimeout: config.WriteTimeout * time.Second,
		MinIdleConns: config.MinIdleConns,
	})
}

func GetRedis() *redis.Client {
	return rdb
}
