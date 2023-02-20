package batch

import (
	"fmt"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
)

type Test1 struct {
}

func (t Test1) RegRouter(engine *web.Engine) {
	g := engine.Group("/test1", func(context *web.Context) {
		fmt.Println("before1")
		context.Next()
		fmt.Println("after1")
	})
	g.GET("/", false, getRedis)
	g.POST("/hello", false, Hello)
}
