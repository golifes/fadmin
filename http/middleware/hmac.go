package middleware

import (
	"bytes"
	"encoding/json"
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"fadmin/pkg/jwt"
	"fadmin/tools/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	uid time params()
"[%--`~!@#$^&*()=|{}':;',\\[\\].<>/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？]"
*/

func Hmac(limit float64) gin.HandlerFunc {
	return func(ctx app.GContext) {
		g := app.G{Context: ctx}

		uid := g.Request.Header.Get("uid")
		uri := g.Request.URL.Path
		method := g.Request.Method
		sign := g.Request.Header.Get("sign")
		timestamp := g.Request.Header.Get("timestamp")
		if uid == "" && sign == "" {
			g.Json(http.StatusForbidden, e.Forbid, "")
			return
		}
		u, err := time.Parse("2006-01-02 15:04:05", timestamp)
		t := time.Now()
		if err != nil && t.Sub(u).Seconds() > limit {
			g.Json(http.StatusForbidden, e.Forbid, "")
			return
		}
		var body []byte
		var tmp []byte
		keys := []string{uid, method, uri, timestamp}
		body, _ = ioutil.ReadAll(g.Request.Body)
		tmp = body
		if tmp != nil && len(tmp) != 0 {
			m := make(map[string]interface{})
			if err = json.Unmarshal([]byte(string(tmp)), &m); err != nil {
				g.Json(http.StatusOK, e.ParamError, "")
				return
			}
			if utils.Xss(m) {
				g.Json(http.StatusOK, e.ParamError, "")
				return
			}
			//给参数复制
			if pn, ok := m["pn"]; ok {
				g.Set("pn", pn)
			} else {
				g.Set("pn", 0)
			}
		}
		if !jwt.Hmac(keys, tmp, sign) {
			g.Json(http.StatusForbidden, e.Forbid, "")
			return
		}
		//这里校验成功
		g.Set("authorized", true)
		g.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		g.Next()
	}
}
