package web

import (
	"github.com/lristar/go-boot/pkg/jaeger"
	"github.com/lristar/go-boot/pkg/sentry"
)

// Config 参数必须大写 要不然viper解析成结构体的时候不识别
type Config struct {
	ServerKey      string              `json:"server_key"`
	ServerName     string              `json:"server_name"`
	Port           int                 `json:"port"`
	LoginAPIPublic string              `json:"login_api_public"`
	UserAPI        string              `json:"user_api"`
	Jaeger         jaeger.JaegerConfig `json:"jaeger"`
	Sentry         sentry.SentryConfig `json:"sentry"`
}
