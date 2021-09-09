package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"service-all/app/init/config"
	"service-all/app/init/global"
	"time"
)

// 声明一个全局的rdb变量

// 初始化连接
func InitRedisClient() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Config.Redis.Host, config.Config.Redis.Port),
		Password: config.Config.Redis.Password, // no password set
		DB:       0,                            // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.LOG.Panic("redis连接失败", zap.Error(err))
	}
	return rdb
}
