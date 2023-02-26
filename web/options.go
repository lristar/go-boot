package web

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"io"
)

type IOption interface {
	Start(string) (io.Closer, error)
}

type Option func(ops *Options) (io.Closer, error)

// Options for web service.
type Options struct {
	beforeStart []func() error
	beforeStop  []func() error
	afterStart  []func() error
	afterStop   []func() error
	//  校验器
	vl  *validator.Validate
	tra ut.Translator
}

func Validate(vl *validator.Validate, tra ut.Translator) {
	application.Options.vl = vl
	application.Options.tra = tra
}

//func (o *Options) init() (io.Closer, error) {
//	jaegerCloser, err := jaeger.InitJaeger(o.serverKey, o.jaegerAddressCollectorEndpoint)
//	if err != nil {
//	}
//	err := sentry.InitSentry(o.sentryUrl, o.serverKey)
//	return jaegerCloser, nil
//}
