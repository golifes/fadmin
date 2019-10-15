package adminc

import (
	"errors"
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	"net/http"
)

func (h HttpAdminHandler) common(ctx app.GContext, obj interface{}) (app.G, error) {
	g := app.G{Context: ctx}

	code := e.Success
	bindJSON := ctx.ShouldBindJSON(&obj)
	if !utils.CheckError(bindJSON, "bind params error") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return g, errors.New("参数绑定失败")
	}
	return g, nil
}
