package example

import (
	"fmt"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	redis "gitlab.gf.com.cn/hk-common/go-tool/server/redis"
	"sync"
)

var (
	rLock sync.RWMutex
)

// Settings Config必须大写
type Settings struct {
	web.Config
	Redis redis.RedisConfig `json:"redis"`
	// 扩展配置
}

func (c *Settings) TryLock() error {
	fmt.Println("加锁")
	if rLock.TryLock() {
		return nil
	}
	return fmt.Errorf("加锁失败")
}

func (c *Settings) UnLock() error {
	fmt.Println("解锁")
	rLock.Unlock()
	return nil
}

func (c *Settings) Get() Settings {
	rLock.RLock()
	defer rLock.RUnlock()
	return *c
}
