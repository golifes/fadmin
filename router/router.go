package router

import (
	_ "fadmin/cmd/docs"
	a "fadmin/http/controller/admin"
	"fadmin/http/middleware"
	"fadmin/pkg/config"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(path string) *Engine {
	r := &Engine{gin.New(), a.NewAdminHttpAdminHandler(path),
		func() string { return config.NewHttpPort() }}
	r.Use(middleware.DummyMiddleware())
	admin(r)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Engine.Run(r.Port())
	return r
}
