package router

import (
	"fadmin/http/controller/admin"
	"fadmin/http/controller/wx"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
	*adminc.HttpAdminHandler
	*wx.HttpWxHandler
	Port func() string

	//*wx.HttpWxHandler
}
