package redis

import (
	"github.com/lristar/go-boot/web"
	redis "github.com/lristar/go-tool/server/redis"
	"io"
)

func InitRedis(config redis.RedisConfig) func(ops *web.Options) (io.Closer, error) {
	return func(ops *web.Options) (io.Closer, error) {
		redis.InitRedisClient(config)
		return redis.NewClient(), nil
	}
}
