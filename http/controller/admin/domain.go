package adminc

import (
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	"net/http"
)

/**
校验域和应用是否存在
*/
func (h HttpAdminHandler) ExistDomainApp(g app.G, ctx app.GContext, did, aid string, model interface{}) int {
	values := []interface{}{did, aid}
	fields := []string{"did=", "aid="}

	count, err := h.logic.Count(g.NewContext(ctx), "", fields, values, model)
	if !utils.CheckError(err, count) || count == 0 {
		g.Json(http.StatusOK, e.ParamError, "")
		return e.ParamError
	}
	return e.Success
}
