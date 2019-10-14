package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func admin(e *Engine) {
	r := e.Group("/user/admin")
	{
		r.GET("/test", func(c *gin.Context) {
			all, err := ioutil.ReadAll(c.Request.Body)
			fmt.Println(string(all), err)
			c.JSON(http.StatusOK, "hello world")
		})
		r.POST("/register", e.Register)
		r.POST("/login", e.Login)
	}
}
