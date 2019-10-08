package router

import (
	adminc "fadmin/http/controller/admin"
	"fadmin/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(path string) Engine {
	r := Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path), func() string {

		return config.NewHttpPort()
	}}
	u := r.Group("/user")
	admin(u)
	err := r.Engine.Run(r.Port())
	fmt.Println(err)
	return r
}
