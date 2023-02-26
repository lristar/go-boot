package mongo

import (
	"github.com/lristar/go-boot/web"
	mg "github.com/lristar/go-tool/server/mongo"
	"io"
)

func InitMg(config mg.Config) func(ops *web.Options) (io.Closer, error) {
	return func(ops *web.Options) (io.Closer, error) {
		mg.InitMongodb(config)
		return mg.GetDefaultServer(), nil
	}
}
