package app

import (
	"context"
	"fadmin/pkg/e"
	"github.com/gin-gonic/gin"
)

type GContext = *gin.Context

type G struct {
	*gin.Context
}

// NewContext 封装上线文入口
func (*G) NewContext(c *gin.Context) context.Context {
	//这里可以做其他事情
	parent := context.Background()
	return parent
}

func (g *G) Json(httpCode, code int, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = e.GetMsg(code)
	m["data"] = data
	g.JSON(httpCode, m)
	return
}
