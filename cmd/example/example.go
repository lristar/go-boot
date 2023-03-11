package main

import (
	"flag"
	"github.com/lristar/go-boot/api/example"
	"github.com/lristar/go-boot/pkg/validator"
	"github.com/lristar/go-boot/web"
	config "github.com/lristar/go-tool/config"
)

var (
	defaultConfigPath *string
	_cf               *example.Settings
)

func init() {
	defaultConfigPath = flag.String("f", ".", "配置文件位置")
	flag.Parse()
}

func main() {
	_cf = new(example.Settings)
	var err error
	if err = config.Setup(*defaultConfigPath, _cf, config.ResetTag("json")); err != nil {
		panic(err)
	}
	web.NewApp(
		_cf,
		// 校验器和翻译器的创建
		validator.InitValidate(),
		//// 开启redis连接
		//redis.InitRedis(_cf.Redis),
		//// 开启pg连接
		//pg.InitPg(_cf.Pg, true),
		//// 开启mongodb连接
		//mongo.InitMg(_cf.Mg),
	).
		UseRoutes(
			example.Routers...,
		).UseMiddleware().
		Run()
}
