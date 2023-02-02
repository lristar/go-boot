package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
)

type Test struct {
}

func (t Test) RegRouter(engine *web.Engine) {
	g := engine.Group("/test", func(context *gin.Context) {
		fmt.Println("before")
		context.Next()
		fmt.Println("after")
	})
	g.GET("/", true, func(c *web.Context) {
		c.JsonOK("ok")
	})
}

func main() {
	web.NewApp(
		"test",
		"测试",
		web.SentryUrl("http://964d306156fe45ddad42b725ade8d247:e79757b5d2f94d839439d440b7096399@10.68.41.33:9000/2"),
		web.JaegerAddressCollectorEndpoint("http://10.68.41.33:8998/api/traces"),
		web.LoginAPIPublic("http://10.68.41.36:9000"),
		web.UserAPI("http://10.68.41.32:8500/api"),
	).
		UseRoutes(
			Test{},
		).
		Run(":8080")
}
