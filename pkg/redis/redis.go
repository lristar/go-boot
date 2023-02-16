package redis

import (
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	redis "gitlab.gf.com.cn/hk-common/go-tool/server/redis"
	"io"
)

func InitRedis(config redis.RedisConfig) func(ops *web.Options) (io.Closer, error) {
	return func(ops *web.Options) (io.Closer, error) {
		redis.InitRedisClient(config)
		return redis.NewClient(), nil
	}
}
