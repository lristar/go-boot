package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lristar/go-validator/validator"

	//"github.com/go-playground/validator/v10"
	"github.com/lristar/go-boot/dto/base"
	myerror "github.com/lristar/go-boot/lib/error"
	isp "github.com/lristar/go-boot/third_api/isp"
	"net/http"
)

// Context 自定义上下文
type Context struct {
	*gin.Context
	User isp.LoginInfo
	App  *Application
	// 扩展配置
}

type HandleFunc func(c *Context)

func NewContext(context *gin.Context) *Context {
	return &Context{Context: context, App: &application}
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

func (c *Context) Validator(s interface{}) error {
	return validator.ValidatorStruct(s)
	//if c.App.vl == nil {
	//	return errors.New("请先注册校验器！")
	//}
	//err := c.App.vl.Struct(s)
	//sType := reflect.TypeOf(s)
	//if sType.Kind() == reflect.Ptr {
	//	sType = sType.Elem()
	//}
	//if err != nil {
	//	msg := ""
	//	if rErr, ok := err.(validator.ValidationErrors); ok {
	//		for _, e := range rErr {
	//			ss := strings.Split(e.StructNamespace(), ".")
	//			ss = ss[1:]
	//			jsonKey := strings.ReplaceAll(stringx.Snake(strings.Join(ss, ".")), "._", ".")
	//			result := e.Translate(c.App.tra)
	//			results := strings.Split(result, " ")
	//			if len(results) > 0 {
	//				msg += strings.Replace(result+";", results[0], jsonKey, 1)
	//			} else {
	//				msg += result + ";"
	//			}
	//		}
	//	}
	//	return errors.New(msg)
	//}
	//return nil
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
