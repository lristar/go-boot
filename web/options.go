package web

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gitlab.gf.com.cn/hk-common/go-boot/jaeger"
	"gitlab.gf.com.cn/hk-common/go-boot/sentry"
	"io"
)

type IOption interface {
	Start(string) error
	io.Closer
}

// Options for web service.
type Options struct {
	// Auth   auth.Auth
	// Broker broker.Broker
	// Before and After funcs
	beforeStart []func() error
	beforeStop  []func() error
	afterStart  []func() error
	afterStop   []func() error
	//  校验器
	vl  *validator.Validate
	tra ut.Translator

	// Other options for implementations of the interface
	// can be stored in a context
	serverKey                      string
	serverName                     string
	loginAPIPublic                 string
	userAPI                        string
	jaegerAddressCollectorEndpoint jaeger.JaegerConfig
	sentryUrl                      sentry.SentryConfig
}

func JaegerAddressCollectorEndpoint(s string) Option {
	return func(ops *Options) {
		ops.jaegerAddressCollectorEndpoint = s
	}
}

func LoginAPIPublic(s string) Option {
	return func(ops *Options) {
		ops.loginAPIPublic = s
	}
}

func UserAPI(s string) Option {
	return func(ops *Options) {
		ops.userAPI = s
	}
}

func SentryUrl(s string) Option {
	return func(ops *Options) {
		ops.sentryUrl = s
	}
}

func Validate(vl *validator.Validate, tra ut.Translator) {
	application.Options.vl = vl
	application.Options.tra = tra
}

func (o *Options) init() (io.Closer, error) {
	jaegerCloser, err := jaeger.InitJaeger(o.serverKey, o.jaegerAddressCollectorEndpoint)
	if err != nil {

	}
	err := sentry.InitSentry(o.sentryUrl, o.serverKey)
	return jaegerCloser, nil
}

func (o *Options) startJaeger() {

}
