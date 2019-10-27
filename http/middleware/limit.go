package middleware

import (
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"fadmin/tools/rdx"
	"fadmin/tools/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Limit() gin.HandlerFunc {
	return func(ctx app.GContext) {
		g := app.G{Context: ctx}

		uid := g.Request.Header.Get("uid")
		method := g.Request.Method
		s := utils.StringJoin("limit_", uid, "_", method)
		//从redis取这个key
		if rdx.Get(s) {
			g.Json(http.StatusOK, e.RequestError, nil)
		} else {
			rdx.Set(s)
		}
		//如果是get请求,先校验是否登录,未登录，只能访问前三页数据
		pn := g.GetInt("pn")
		if !g.GetBool("authorized") && pn >= 3 {
			g.Json(http.StatusOK, e.NoLogin, nil)
		}
	}
}
