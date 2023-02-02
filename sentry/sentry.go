package sentry

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/getsentry/sentry-go"
	"gitlab.gf.com.cn/hk-common/go-boot/logger"
)

var serverName string

func InitSentry(sentryUrl, serverKey string) {
	serverName = serverKey
	// sentry 初始化
	e := raven.SetDSN(sentryUrl)
	if e != nil {
		panic("Sentry 启动失败：" + e.Error())
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryUrl,
	}); err != nil {
		panic(fmt.Errorf("Sentry initialization failed: %w\n", err))
	}
}

var tags = map[string]string{"server_name": serverName}

// Panic 执行函数并捕获+报告错误
func Panic(f func()) {
	raven.CapturePanic(f, tags)
}

// Sentry 报告错误
func Sentry(err error, args ...raven.Interface) {
	raven.CaptureError(err, tags, args...)
}

// LogAndSentry 报告并打印错误
func LogAndSentry(e error, args ...raven.Interface) {
	logger.Errorf("%s\n", e.Error())
	Sentry(e, args...)
}
