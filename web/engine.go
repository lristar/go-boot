package web

import (
	"github.com/gin-gonic/gin"
)

// IRegRouter 注册路由
type IRegRouter interface {
	RegRouter(engine *Engine)
}

type Engine struct {
	*gin.Engine
	rg *RouterGroup
}

type Option func(ops *Options)

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
