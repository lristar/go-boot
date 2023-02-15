package main

import (
	"flag"
	"fmt"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	"gitlab.gf.com.cn/hk-common/go-tool/config"
)

var (
	defaultConfigPath *string
	_cf               *web.Config
)

func init() {
	defaultConfigPath = flag.String("f", "./config/", "配置文件位置")
	flag.Parse()
}

type Test1 struct {
}

func (t Test1) RegRouter(engine *web.Engine) {
	g := engine.Group("/test1", func(context *web.Context) {
		fmt.Println("before1")
		context.Next()
		fmt.Println("after1")
	})
	g.GET("/", false, func(c *web.Context) {
		c.JsonOK(_cf.Get().ServerName)
	})
}

type Test2 struct {
}

func (t Test2) RegRouter(engine *web.Engine) {
	g := engine.Group("/test2", func(context *web.Context) {
		fmt.Println("before2")
		context.Next()
		fmt.Println("after2")
	})
	g.GET("/", true, func(c *web.Context) {
		c.JsonOK("ok")
	})
}

func main() {
	_cf = new(web.Config)
	var err error
	if err = config.Setup(*defaultConfigPath, _cf, config.ResetTag("json")); err != nil {
		panic(err)
	}
	web.NewApp(
		_cf.Get(),
	).
		UseRoutes(
			Test1{},
			Test2{},
		).UseMiddleware().
		Run()
}
