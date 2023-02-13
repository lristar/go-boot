package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.gf.com.cn/hk-common/go-boot/middleware"
	"net/http"
)

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
