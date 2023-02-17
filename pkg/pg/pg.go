package pg

import (
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	pg "gitlab.gf.com.cn/hk-common/go-tool/server/pg"
	"io"
)

func InitPg(config pg.Config, debug bool) func(ops *web.Options) (io.Closer, error) {
	return func(ops *web.Options) (io.Closer, error) {
		pg.InitPg(config)
		return pg.GetDb(debug), nil
	}
}
