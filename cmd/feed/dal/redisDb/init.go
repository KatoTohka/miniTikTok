package redisDb

import (
	"github.com/go-redis/redis/v8"
	"miniTikTok/pkg/constants"
)

var RDB *redis.Client

func Init() {
	//初始化redis，连接地址和端口，密码，数据库名称
	RDB = redis.NewClient(&redis.Options{
		Addr:     constants.RedisDefaultAddr,
		Password: constants.RedisDefaultPassWord,
		DB:       constants.RedisDefaultDB,
	})
}
