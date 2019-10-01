package router

import (
	"fadmin/http/controller/admin"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
	*adminc.HttpAdminHandler
	Port func() string
	//*wx.HttpWxHandler
}
