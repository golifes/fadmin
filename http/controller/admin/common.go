package adminc

import (
	"errors"
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	"net/http"
)

func (h HttpAdminHandler) common(ctx app.GContext, model interface{}) (app.G, error) {
	g := app.G{Context: ctx}

	code := e.Success
	if !utils.CheckError(ctx.ShouldBindJSON(&model), "login") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return g, errors.New("参数绑定失败")
	}
	return g, nil
}
