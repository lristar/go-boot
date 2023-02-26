package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lristar/go-boot/middleware"
	"net/http"
)

// IRegRouter 注册路由
type IRegRouter interface {
	RegRouter(engine *Engine)
}

type Engine struct {
	*gin.Engine
	rg *RouterGroup
}

func (service *Engine) use(middleware ...HandleFunc) *Engine {
	arr := make([]gin.HandlerFunc, len(middleware))
	for i, f := range middleware {
		arr[i] = Handle(f)
	}
	service.Engine.Use(arr...)
	return service
}

func (service *Engine) Group(relativePath string, handlers ...HandleFunc) *RouterGroup {
	arr := make([]gin.HandlerFunc, len(handlers))
	for i, f := range handlers {
		arr[i] = Handle(f)
	}
	rg := service.Engine.Group(relativePath, arr...)
	service.rg.RouterGroup = rg
	return service.rg
}

type RouterGroup struct {
	*gin.RouterGroup
}

func (rg *RouterGroup) handle(httpMethod string, checkLogin bool, relativePath string, handlers ...HandleFunc) {
	arr := make([]gin.HandlerFunc, 0)
	if checkLogin {
		arr = append(arr, middleware.MCheckLogin(application.cf.ServerKey, application.cf.LoginAPIPublic, application.cf.UserAPI))
	}
	for _, handler := range handlers {
		arr = append(arr, Handle(handler))
	}
	rg.Handle(httpMethod, relativePath, arr...)
}

func (rg *RouterGroup) POST(relativePath string, checkLogin bool, handlers ...HandleFunc) {
	rg.handle(http.MethodPost, checkLogin, relativePath, handlers...)
}

func (rg *RouterGroup) GET(relativePath string, checkLogin bool, handlers ...HandleFunc) {
	rg.handle(http.MethodGet, checkLogin, relativePath, handlers...)
}

func (rg *RouterGroup) PUT(relativePath string, checkLogin bool, handlers ...HandleFunc) {
	rg.handle(http.MethodPut, checkLogin, relativePath, handlers...)
}

func (rg *RouterGroup) DELETE(relativePath string, checkLogin bool, handlers ...HandleFunc) {
	rg.handle(http.MethodDelete, checkLogin, relativePath, handlers...)
}
