package router

import (
	adminc "fadmin/http/controller/admin"
	"fadmin/http/middleware"
	"fadmin/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(path string) *Engine {
	r := &Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path), func() string {

		return config.NewHttpPort()
	}}
	r.Use(middleware.DummyMiddleware())
	//u := r.Group("/user")
	admin(r)
	err := r.Engine.Run(r.Port())
	fmt.Println(err)
	return r
}
