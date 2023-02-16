package batch

import (
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	myredis "gitlab.gf.com.cn/hk-common/go-tool/server/redis"
)

func Hello(c *web.Context) {
	cli := myredis.NewClient()
	if err := cli.Set(c, "hahah", "test111", 0).Err(); err != nil {
		c.JsonError(err)
	}
	c.JsonOK("ok")
}
