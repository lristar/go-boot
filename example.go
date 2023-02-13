package main

import (
	"flag"
	"fmt"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	"gitlab.gf.com.cn/hk-common/go-tool/config"
)

var (
	configName = flag.String("f", "./config/config.yaml", "配置文件位置")
)

type Test1 struct {
}

func (t Test1) RegRouter(engine *web.Engine) {
	g := engine.Group("/test1", func(context *web.Context) {
		fmt.Println("before1")
		context.Next()
		fmt.Println("after1")
	})
	g.GET("/", false, func(c *web.Context) {
		c.JsonOK("ok")
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
	var cf *web.Settings
	var err error
	if err = config.Setup(*configName, &cf, config.ResetTag("yaml")); err != nil {
		panic(err)
	}
	web.NewApp(
		cf.Config,
	).
		UseRoutes(
			Test1{},
			Test2{},
		).UseMiddleware().
		Run()
}
