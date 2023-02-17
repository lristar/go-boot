package main

import (
	"flag"
	"gitlab.gf.com.cn/hk-common/go-boot/api/example"
	"gitlab.gf.com.cn/hk-common/go-boot/pkg/mongo"
	"gitlab.gf.com.cn/hk-common/go-boot/pkg/pg"
	"gitlab.gf.com.cn/hk-common/go-boot/pkg/redis"
	"gitlab.gf.com.cn/hk-common/go-boot/pkg/validator"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	"gitlab.gf.com.cn/hk-common/go-tool/config"
)

var (
	defaultConfigPath *string
	_cf               *example.Settings
)

func init() {
	defaultConfigPath = flag.String("f", "./cmd/example", "配置文件位置")
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
		// 校验器和翻译器的创建
		validator.InitValidate(),
		// 开启redis连接
		redis.InitRedis(_cf.Redis),
		// 开启pg连接
		pg.InitPg(_cf.Pg, true),
		// 开启mongodb连接
		mongo.InitMg(_cf.Mg),
	).
		UseRoutes(
			example.Routers...,
		).UseMiddleware().
		Run()
}
