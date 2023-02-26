package batch

import (
	"github.com/lristar/go-boot/web"
	myredis "github.com/lristar/go-tool/server/redis"
	"github.com/lristar/go-validator/validator"
)

func getRedis(c *web.Context) {
	cli := myredis.NewClient()
	if err := cli.Set(c, "hahah", "test111", 0).Err(); err != nil {
		c.JsonError(err)
	}
	rs, err := cli.Get(c, "hahah").Result()
	if err != nil {
		c.JsonError(err)
	}

	c.JsonOK(rs)
}

type HelloReq struct {
	Name string `json:"name" validate:"hello"`
}

func Hello(c *web.Context) {
	req := new(HelloReq)
	if err := validator.ValidatorStruct(req); err != nil {
		c.JsonError(err)
		return
	}
	c.JsonOK("hahah")
}
