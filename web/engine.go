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

func (service *Engine) Use(middleware ...gin.HandlerFunc) *Engine {
	service.Engine.Use(middleware...)
	return service
}

func (service *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	rg := service.Engine.Group(relativePath, handlers...)
	service.rg.RouterGroup = rg
	return service.rg
}
