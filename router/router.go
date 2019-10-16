package router

import (
	a "fadmin/http/controller/admin"
	"fadmin/http/middleware"
	"fadmin/pkg/config"
	"github.com/gin-gonic/gin"
)

func InitRouter(path string) *Engine {
	r := &Engine{gin.New(), a.NewAdminHttpAdminHandler(path),
		func() string { return config.NewHttpPort() }}
	r.Use(middleware.DummyMiddleware())
	admin(r)
	//r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Engine.Run(r.Port())
	return r
}
