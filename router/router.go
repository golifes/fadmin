package router

import (
	_ "fadmin/cmd/docs"
	a "fadmin/http/controller/admin"
	"fadmin/http/middleware"
	"fadmin/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"net/http"
)

func InitRouter(path string) *Engine {
	r := &Engine{gin.New(), a.NewAdminHttpAdminHandler(path),
		func() string { return config.NewHttpPort() }}
	r.Use(middleware.DummyMiddleware())
	//admin(r)
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/test", demo)
	r.Engine.Run(r.Port())
	return r
}

//CreatScene createScene
// @Summary createScene
// @Description createScene
// @Accept multipart/form-data
// @Produce  json
// @Param app_key formData string true "AppKey"
// @Param nonce_str formData string true "NonceStr"
// @Param time_stamp formData string true "TimeStamp"
// @Success 200 {object} app.R
// @Failure 500 {object} app.R
// @Router /dictionaries/createScene [post]
func demo(c *gin.Context) {

	all, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(all), err)
	c.JSON(http.StatusOK, "hello world")
}
