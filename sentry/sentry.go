package sentry

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/getsentry/sentry-go"
	"gitlab.gf.com.cn/hk-common/go-boot/lib/atomic"
	"gitlab.gf.com.cn/hk-common/go-boot/logger"
	"io"
)

var (
	enabled    atomic.AtomicBool
	closer     io.Closer
	serverName string
)

type SentryConfig struct {
	sentryUrl string
}

func (s *SentryConfig) Start(serverKey string) error {
	serverName = serverKey
	// sentry 初始化
	err := raven.SetDSN(s.sentryUrl)
	if err != nil {
		return err
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: s.sentryUrl,
	}); err != nil {
		return fmt.Errorf("Sentry initialization failed: %w\n", err)
	}
	enabled.Set(true)
	return nil
}

func (s *SentryConfig) Close() error {
	if closer != nil {
		return closer.Close()
	}
	return nil
}

func InitSentry(sentryUrl, serverKey string) error {
	serverName = serverKey
	// sentry 初始化
	e := raven.SetDSN(sentryUrl)
	if e != nil {
		panic("Sentry 启动失败：" + e.Error())
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryUrl,
	}); err != nil {
		return fmt.Errorf("Sentry initialization failed: %w\n", err)
	}
	return nil
}

var tags = map[string]string{"server_name": serverName}

// Panic 执行函数并捕获+报告错误
func Panic(f func()) {
	if enabled.True() {
		raven.CapturePanic(f, tags)
	}
}

// Sentry 报告错误
func Sentry(err error, args ...raven.Interface) {
	if enabled.True() {
		raven.CaptureError(err, tags, args...)
	}
}

// LogAndSentry 报告并打印错误
func LogAndSentry(e error, args ...raven.Interface) {
	logger.Errorf("%s\n", e.Error())
	Sentry(e, args...)
}
