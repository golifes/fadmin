package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func admin(r *gin.RouterGroup) {
	u := r.Group("/admin")
	{
		u.GET("/test", func(c *gin.Context) {
			all, err := ioutil.ReadAll(c.Request.Body)
			fmt.Println(string(all), err)
			c.JSON(http.StatusOK, "hello world")
		})
	}
}
