package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	"gobase/config"
)

type RedisConnector struct {
	config.RedisConf
	RedisClient *redis.Client
}

var redisConnector *RedisConnector

func GetRedisConnector() *RedisConnector {
	appConfig := config.GetConfig()
	if redisConnector == nil {
		redisConnector = &RedisConnector{
			RedisConf: config.RedisConf{
				Host:     appConfig.Redis.Host,
				Port:     appConfig.Redis.Port,
				Username: appConfig.Redis.Username,
				Password: appConfig.Redis.Password,
			},
		}
	}
	return redisConnector
}

func (conn *RedisConnector) Connect() *redis.Client {
	appConfig := config.GetConfig()
	conn.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", appConfig.Redis.Host, appConfig.Redis.Port),
		Password: appConfig.Redis.Password,
		DB:       0,
	})
	return conn.RedisClient
}

func GetRedisClient() *redis.Client {
	return redisConnector.RedisClient
}
