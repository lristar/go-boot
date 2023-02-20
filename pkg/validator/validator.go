package validator

import (
	"github.com/go-playground/validator/v10"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	val "gitlab.gf.com.cn/hk-common/validator/validator"
	"io"
)

type Validator struct{}

func (v *Validator) Close() error {
	return nil
}

func InitValidate() func(ops *web.Options) (io.Closer, error) {
	return func(ops *web.Options) (io.Closer, error) {
		val.InitValidator()
		// 自定义校验器列表
		if err := val.AddRegisterVal("hello", Hello, true); err != nil {
			return nil, err
		}
		return new(Validator), nil
	}
}

func Hello(fl validator.FieldLevel) bool {
	if fl.Field().String() == "hello" {
		return true
	}
	return false
}
