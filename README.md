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
	_cf = new(example.Settings)
	var err error
	if err = config.Setup(*defaultConfigPath, _cf, config.ResetTag("json")); err != nil {
		panic(err)
	}
	web.NewApp(
		_cf.Get().Config,
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


```

And use the Go command to run the demo:

```
# run example.go and visit 0.0.0.0:8080/ping on browser
$ go run cmd/example/example.go
```
