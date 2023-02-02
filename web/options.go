package web

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gitlab.gf.com.cn/hk-common/go-boot/jaeger"
	"gitlab.gf.com.cn/hk-common/go-boot/sentry"
	"io"
)

// Options for web service.
type Options struct {
	//Auth   auth.Auth
	//Broker broker.Broker
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
	jaegerAddressCollectorEndpoint string
	sentryUrl                      string
	loginAPIPublic                 string
	userAPI                        string
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

var jaegerCloser io.Closer

func (o Options) init() error {
	jaegerCloser = jaeger.InitJaeger(o.serverKey, o.jaegerAddressCollectorEndpoint)
	sentry.InitSentry(o.sentryUrl, o.serverKey)
	return nil
}
