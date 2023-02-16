package main

import (
	"flag"
	"gitlab.gf.com.cn/hk-common/go-boot/api/example"
	"gitlab.gf.com.cn/hk-common/go-boot/pkg/redis"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	"gitlab.gf.com.cn/hk-common/go-tool/config"
)

var (
	defaultConfigPath *string
	_cf               *example.Settings
)

func init() {
	defaultConfigPath = flag.String("f", "./config/", "配置文件位置")
	flag.Parse()
}

func main() {
	_cf = new(example.Settings)
	var err error
	if err = config.Setup(*defaultConfigPath, _cf, config.ResetTag("json")); err != nil {
		panic(err)
	}
	web.NewApp(
		_cf.Get().Config,
		redis.InitRedis(_cf.Redis),
	).
		UseRoutes(
			example.Routers...,
		).UseMiddleware().
		Run()
}
