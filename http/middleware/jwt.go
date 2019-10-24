package middleware

import (
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"fadmin/pkg/jwt"
	"github.com/gin-gonic/gin"

	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c app.GContext) {
		var code int
		g := app.G{c}
		code = e.Success
		/**
		header中带用户id
		*/
		token := c.Request.Header.Get("token")
		//先查询redis存在不  key是用户id
		/**
		if userKey {

		}else{

		}
		*/
		if token == "" {
			code = e.Unauthorized
		} else {
			//根据实际需要设置数据
			if claims, code := jwt.ParseToken(token); code == e.Success {
				c.Set("userName", claims.Username)
				c.Set("userId", claims.Id)
				c.Set("isAdmin", claims.IsAdmin)
				c.Set("isRoot", claims.IsRoot)
			}
		}
		if code != e.Success {
			g.Json(http.StatusOK, code, "")
			g.Abort()
			return
		}
		g.Next()
	}
}
