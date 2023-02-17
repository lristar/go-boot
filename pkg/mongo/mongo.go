package mongo

import (
	"gitlab.gf.com.cn/hk-common/go-boot/web"
	mg "gitlab.gf.com.cn/hk-common/go-tool/server/mongo"
	"io"
)

func InitMg(config mg.Config) func(ops *web.Options) (io.Closer, error) {
	return func(ops *web.Options) (io.Closer, error) {
		mg.InitMongodb(config)
		return mg.GetDefaultServer(), nil
	}
}
