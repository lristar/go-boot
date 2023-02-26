package user

import (
	"fmt"
	"github.com/lristar/go-boot/web"
)

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
