package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.gf.com.cn/hk-common/go-boot/logger"
	"gitlab.gf.com.cn/hk-common/go-boot/middleware"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var application Application

type RegistryFunc func()

type Application struct {
	*Engine
	unRegistryFunc RegistryFunc
	*Options
}

type IRegistry interface {
	Reg()
}

type IUnRegistry interface {
	UnReg()
}

// NewApp 实例化应用
func NewApp(serverKey, serverName string, opts ...Option) *Application {
	e := gin.New()
	application = Application{
		Engine: &Engine{
			Engine: e,
			rg: &RouterGroup{
				&e.RouterGroup,
			},
		},
		Options: &Options{
			beforeStart: nil,
			beforeStop:  nil,
			afterStart:  nil,
			afterStop:   nil,
			serverKey:   serverKey,
			serverName:  serverName,
		},
	}

	// options赋值
	for _, opt := range opts {
		opt(application.Options)
	}
	if err := application.init(); err != nil {
		panic(err)
	}

	// 注册基础中间件
	e.Use(gin.Recovery(), gin.Logger(), middleware.PrintBody())

	// 注册jaeger
	if jaegerCloser != nil {
		defer jaegerCloser.Close()
		e.Use(middleware.Jaeger())
	}

	// 注册sentry
	e.Use(middleware.MSentry(application.serverKey))
	return &application
}

// UseMiddleware 注册中间件
func (app *Application) UseMiddleware(middleware ...gin.HandlerFunc) *Application {
	app.Engine.use(middleware...)
	return app
}

// UseRoutes UseRoutes 注册路由
func (app *Application) UseRoutes(controllers ...IRegRouter) *Application {
	for _, controller := range controllers {
		controller.RegRouter(app.Engine)
	}
	return app
}

// Register 服务注册
func (app *Application) Register(reg RegistryFunc) *Application {
	// 服务注册
	reg()
	return app
}

// Deregister 服务反注册
func (app *Application) Deregister(f RegistryFunc) *Application {
	app.unRegistryFunc = f
	return app
}

// Run 启动应用
func (app *Application) Run(addr ...string) {
	go func(app *Application) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			s := <-c
			logger.Infof("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				logger.Info("application exit")
				// 服务反注册
				if app.unRegistryFunc != nil {
					app.unRegistryFunc()
				}
				time.Sleep(time.Second)
				os.Exit(0)
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}(app)
	err := app.Engine.Run(addr...)
	if err != nil {
		panic(err)
	}
}
