package example

import (
	"gitlab.gf.com.cn/hk-common/go-boot/api/example/service/batch"
	"gitlab.gf.com.cn/hk-common/go-boot/api/example/service/user"
	"gitlab.gf.com.cn/hk-common/go-boot/web"
)

var Routers = []web.IRegRouter{
	batch.Test1{},
	user.Test2{},
}
