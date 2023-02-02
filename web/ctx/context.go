package ctx

import (
	"github.com/gin-gonic/gin"
	"gitlab.gf.com.cn/hk-common/go-boot/dto/base"
	"gitlab.gf.com.cn/hk-common/go-boot/isp"
	myerror "gitlab.gf.com.cn/hk-common/go-boot/lib/error"
	"net/http"
)

// Context 自定义上下文
type Context struct {
	*gin.Context
	User isp.LoginInfo
}

func NewContext(context *gin.Context) *Context {
	return &Context{Context: context}
}

// Handle Context转*gin.Context
func Handle(f func(*Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := NewContext(c)
		context.Context = c
		if user, ok := c.Get("user"); ok {
			context.User = user.(isp.LoginInfo)
		}
		f(context)
	}
}

func (c *Context) JsonOK(res interface{}) {
	c.JsonOKWithStatusCode(http.StatusOK, res)
}

func (c *Context) JsonOKWithStatusCode(code int, res interface{}) {
	r := base.Result{
		ErrCode: 0,
		ErrMsg:  "",
		Data:    res,
	}
	c.JSON(code, r)
}

func (c *Context) JsonError(err error) {
	c.JsonErrorWithStatusCode(http.StatusBadRequest, err)
}

func (c *Context) JsonErrorWithStatusCode(code int, err error) {
	res := base.Result{}
	switch err.(type) {
	case myerror.IError:
		e := err.(myerror.IError)
		res.ErrCode = e.Code()
		res.ErrMsg = e.Error()
	case error:
		res.ErrMsg = err.Error()
	default:
		res.ErrMsg = err.Error()
	}
	c.JSON(code, res)
}
