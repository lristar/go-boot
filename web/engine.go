package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.gf.com.cn/hk-common/go-boot/middleware"
	"gitlab.gf.com.cn/hk-common/go-boot/web/ctx"
	"net/http"
)

// IRegRouter 注册路由
type IRegRouter interface {
	RegRouter(engine *Engine)
}

type Engine struct {
	*gin.Engine
}

type Option func(ops *Options)

type handleFunc func(c *ctx.Context)

func (service *Engine) handle(httpMethod string, checkLogin bool, relativePath string, handlers ...handleFunc) {
	arr := make([]gin.HandlerFunc, 0)
	if checkLogin {
		arr = append(arr, ctx.Handle(middleware.MCheckLogin(application.serverKey, application.loginAPIPublic, application.userAPI)))
	}
	for _, handler := range handlers {
		arr = append(arr, ctx.Handle(handler))
	}
	service.Handle(httpMethod, relativePath, arr...)
}

func (service *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *Engine {
	rg := service.Engine.Group(relativePath, handlers...)
	service.Engine.RouterGroup = *rg
	return service
}

func (service *Engine) Use(middleware ...gin.HandlerFunc) *Engine {
	service.Engine.Use(middleware...)
	return service
}

func (service *Engine) POST(relativePath string, checkLogin bool, handlers ...handleFunc) {
	service.handle(http.MethodPost, checkLogin, relativePath, handlers...)
}

func (service *Engine) GET(relativePath string, checkLogin bool, handlers ...handleFunc) {
	service.handle(http.MethodGet, checkLogin, relativePath, handlers...)
}

func (service *Engine) PUT(relativePath string, checkLogin bool, handlers ...handleFunc) {
	service.handle(http.MethodPut, checkLogin, relativePath, handlers...)
}

func (service *Engine) DELETE(relativePath string, checkLogin bool, handlers ...handleFunc) {
	service.handle(http.MethodDelete, checkLogin, relativePath, handlers...)
}
