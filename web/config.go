package web

import (
	"fmt"
	"gitlab.gf.com.cn/hk-common/go-boot/jaeger"
	"gitlab.gf.com.cn/hk-common/go-boot/sentry"
	"sync"
)

var (
	rLock sync.RWMutex
)

// Settings Config必须大写
type Settings struct {
	Config Config
	// 扩展配置
}

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

func (c *Config) TryLock() error {
	fmt.Println("加锁")
	if rLock.TryLock() {
		return nil
	}
	return fmt.Errorf("加锁失败")
}

func (c *Config) UnLock() error {
	fmt.Println("解锁")
	rLock.Unlock()
	return nil
}

func (c *Config) Get() Config {
	rLock.RLock()
	defer rLock.RUnlock()
	return *c
}
