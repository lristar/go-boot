package example

import (
	"fmt"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	mg "gitlab.gf.com.cn/hk-common/go-tool/server/mongo"
	pg "gitlab.gf.com.cn/hk-common/go-tool/server/pg"
	redis "gitlab.gf.com.cn/hk-common/go-tool/server/redis"
	"sync"
)

var (
	rLock sync.RWMutex
)

// Settings Config必须大写
type Settings struct {
	Config web.Config        `json:"config"`
	Redis  redis.RedisConfig `json:"redis"`
	Pg     pg.Config         `json:"pg"`
	Mg     mg.Config         `json:"mg"`
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
