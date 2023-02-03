
## Getting started

### Getting go-boot

With [Go module] support, simply add the following import

```
import "gitlab.gf.com.cn/hk-common/go-boot"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `go-boot` package:

```sh
$ go get -u gitlab.gf.com.cn/hk-common/go-boot
```

### Running go-boot

First you need to import Gin package for using go-boot, one simplest example likes the follow `example.go`:

```go
package main

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
	web.NewApp(
		"test",
		"测试",
		web.SentryUrl("http://964d306156fe45ddad42b725ade8d247:e79757b5d2f94d839439d440b7096399@10.68.41.33:9000/2"),
		web.JaegerAddressCollectorEndpoint("http://10.68.41.33:8998/api/traces"),
		web.LoginAPIPublic("http://10.68.41.36:9000"),
		web.UserAPI("http://10.68.41.32:8500/api"),
	).
		UseRoutes(
			Test1{},
			Test2{},
		).UseMiddleware().
		Run(":8080")
}


```

And use the Go command to run the demo:

```
# run example.go and visit 0.0.0.0:8080/ping on browser
$ go run example.go
```