package web

import (
	"gitlab.gf.com.cn/hk-common/go-boot/jaeger"
	"gitlab.gf.com.cn/hk-common/go-boot/sentry"
)

// Settings Config必须大写
type Settings struct {
	Config Config
}

// Config 参数必须大写 要不然viper解析成结构体的时候不识别
type Config struct {
	ServerKey      string
	ServerName     string
	Port           int
	LoginAPIPublic string
	UserAPI        string
	Jaeger         jaeger.JaegerConfig
	Sentry         sentry.SentryConfig
}
