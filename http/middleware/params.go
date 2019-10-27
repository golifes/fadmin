package middleware

import (
	"bytes"
	"encoding/json"
	"fadmin/pkg/app"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func DummyMiddleware() gin.HandlerFunc {

	return func(ctx app.GContext) {
		var body []byte
		var tmp []byte

		body, _ = ioutil.ReadAll(ctx.Request.Body)
		tmp = body
		if tmp != nil && len(tmp) != 0 {
			//flag := false
			m := make(map[string]interface{})
			err := json.Unmarshal([]byte(string(tmp)), &m)
			if err != nil {
				ctx.JSON(http.StatusOK, "参数异常")
				ctx.Abort()
			} else {
				for _, v := range m {
					log.Printf("参数值是%s", v)
				}
			}
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		ctx.Next()
	}

}
