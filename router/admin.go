package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func admin(r *gin.RouterGroup) {
	u := r.Group("/admin")
	{
		u.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, "hello world")
		})
	}
}
