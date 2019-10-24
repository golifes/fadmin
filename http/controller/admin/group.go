package adminc

import (
	"fadmin/model/admin"
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"net/http"
)

func (h HttpAdminHandler) AddGroup(ctx app.GContext) {
	var p admin.Group
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	//校验aid did是否存在
	exist := h.ExistDomainApp(g, ctx, p.Did, p.Aid)
	if !exist {
		g.Json(http.StatusOK, e.DomainNotExist, p)
		return
	}
	p.Status = 1

	err = h.logic.TxInsert(g.NewContext(ctx), p)
	if err != nil {
		g.Json(http.StatusOK, e.Errors, p.Name)
	} else {
		g.Json(http.StatusOK, e.Success, p.Name)
	}
}
