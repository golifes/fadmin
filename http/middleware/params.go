package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func DummyMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		var body []byte
		var tmp []byte

		body, _ = ioutil.ReadAll(c.Request.Body)
		tmp = body
		if tmp != nil {
			flag := false
			m := make(map[string]interface{})
			err := json.Unmarshal([]byte(string(tmp)), &m)
			if err != nil {
				c.JSON(http.StatusOK, "参数异常")
				c.Abort()
			} else {
				for _, v := range m {
					fmt.Println(v)
					//这里做字符串参数校验
					if v == "" {
						flag = true
					}
				}
			}

			if flag {
				c.Abort()
			}
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Next()
	}

}
