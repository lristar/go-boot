package web

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gitlab.gf.com.cn/hk-common/go-boot/jaeger"
	"gitlab.gf.com.cn/hk-common/go-boot/sentry"
	"io"
)

type IOption interface {
	Start(string) (io.Closer, error)
}

type Option func(ops *Options) (io.Closer, error)

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
	jaegerAddressCollectorEndpoint *jaeger.JaegerConfig
	sentryUrl                      *sentry.SentryConfig
}

func JaegerAddressCollectorEndpoint(s string) Option {
	return func(ops *Options) (io.Closer, error) {
		ops.jaegerAddressCollectorEndpoint = &jaeger.JaegerConfig{JaegerAddressCollectorEndpoint: s}
		c, err := ops.jaegerAddressCollectorEndpoint.Start(ops.serverKey)
		if err != nil {
			return nil, err
		}
		return c, err
	}
}

func LoginAPIPublic(s string) Option {
	return func(ops *Options) (io.Closer, error) {
		ops.loginAPIPublic = s
		return nil, nil
	}
}

func UserAPI(s string) Option {
	return func(ops *Options) (io.Closer, error) {
		ops.userAPI = s
		return nil, nil
	}
}

func SentryUrl(s string) Option {
	return func(ops *Options) (io.Closer, error) {
		ops.sentryUrl = &sentry.SentryConfig{s}
		c, err := ops.sentryUrl.Start(ops.serverKey)
		if err != nil {
			return nil, err
		}
		return c, err
	}
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
