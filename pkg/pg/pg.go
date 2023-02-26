package pg

import (
	"github.com/lristar/go-boot/web"
	pg "github.com/lristar/go-tool/server/pg"
	"io"
)

func InitPg(config pg.Config, debug bool) func(ops *web.Options) (io.Closer, error) {
	return func(ops *web.Options) (io.Closer, error) {
		pg.InitPg(config)
		return pg.GetDb(debug), nil
	}
}
