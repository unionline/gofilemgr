package env

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gofilemgr/internal/initializers/config"
	"sync"
)

// Redis ...
var Redis *redis.Client
var once = &sync.Once{}

// Init ...
func RedisInit() {
	once.Do(func() {
		options := &redis.Options{
			Addr:     config.Setting.Redis.Addr,
			Password: config.Setting.Redis.Password,
			DB:       config.Setting.Redis.DB,
		}

		Redis = redis.NewClient(options)
		if err := Redis.Ping().Err(); err != nil {
			logrus.Error("redis ping error:", err)
		}
	})
}
