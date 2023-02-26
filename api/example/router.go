package example

import (
	"github.com/lristar/go-boot/api/example/service/batch"
	"github.com/lristar/go-boot/api/example/service/user"
	"github.com/lristar/go-boot/web"
)

var Routers = []web.IRegRouter{
	batch.Test1{},
	user.Test2{},
}
