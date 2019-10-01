package router

import (
	adminc "fadmin/http/controller/admin"
	"github.com/gin-gonic/gin"
)

func InitRouter(path string) Engine {
	r := Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path)}
	u := r.Group("/user")
	admin(u)

	r.Engine.Run(":8080")

	return r
}
