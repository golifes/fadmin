package router

import (
	_ "fadmin/cmd/docs"
	a "fadmin/http/controller/admin"
	"fadmin/http/controller/wx"
	"fadmin/pkg/config"
	"github.com/gin-gonic/gin"
)

func InitRouter(path ...string) *Engine {
	r := &Engine{Engine: gin.New(), Port: func() string { return config.NewHttpPort() }}

	if len(path) == 1 {
		r.HttpAdminHandler = a.NewAdminHttpAdminHandler(path[0])
		r.HttpWxHandler = wx.NewHttpWxHandler(path[0])
	} else if len(path) == 2 {
		r.HttpAdminHandler = a.NewAdminHttpAdminHandler(path[0])
		r.HttpWxHandler = wx.NewHttpWxHandler(path[1])
	} else {
		panic("path params is error")
	}
	//
	//r := &Engine{gin.New(), a.NewAdminHttpAdminHandler(path[0]), wx.NewHttpWxHandler(path[0]),
	//	func() string { return config.NewHttpPort() }}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.DummyMiddleware())
	r.admin()
	r.weiXin()

	//r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Engine.Run(r.Port())
	return r
}
