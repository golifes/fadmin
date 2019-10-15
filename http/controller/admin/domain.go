package adminc

import (
	"fadmin/model/admin"
	"fadmin/pkg/app"
	"fadmin/pkg/config"
	"fadmin/pkg/e"
	"net/http"
)

/**
校验域和应用是否存在
*/
func (h HttpAdminHandler) ExistDomainApp(g app.G, ctx app.GContext, did, aid int64, model interface{}) int {
	//values := []interface{}{did, aid}
	//fields := []string{"did", "aid"}

	//count, err := h.logic.Count(g.NewContext(ctx), "domain_app", fields, values, model)
	//if !utils.CheckError(err, count) || count == 0 {
	//	g.Json(http.StatusOK, e.ParamError, "")
	//	return e.ParamError
	//}
	return e.Success
}

func (h HttpAdminHandler) AddDomain(ctx app.GContext) {
	var p admin.Domain
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	exist := h.logic.Exist(g.NewContext(ctx), &admin.Domain{Name: p.Name})
	if exist {
		g.Json(http.StatusOK, e.DomainExist, "")
		return
	}

	p.Status = 1
	p.Id = config.NewNodeId()
	err = h.logic.TxInsert(g.NewContext(ctx), p)
	if err != nil {
		g.Json(http.StatusOK, e.Errors, "")
	} else {
		g.Json(http.StatusOK, e.Success, "")
	}
}

func (h HttpAdminHandler) DeleteDomain(ctx app.GContext) {
	var p admin.Domain
	//g := app.G{Context: ctx}
	//
	//code := e.Success
	//bindJSON := ctx.ShouldBindJSON(&p)
	//fmt.Println(bindJSON)
	//if !utils.CheckError(bindJSON, "bind params error") {
	//	code = e.ParamError
	//	g.Json(http.StatusOK, code, "")
	//	return
	//}
	g, err := h.common(ctx, p)
	if err != nil {
		return
	}
	//if p.Id == 0 {
	//	g.Json(http.StatusOK, e.ParamError, "")
	//	return
	//}
	//affect, err := h.logic.Delete(g.NewContext(ctx), p.Id, p)
	//if err != nil {
	//	g.Json(http.StatusOK, e.DomainDeleteError, p.Id)
	//	return
	//}
	g.Json(http.StatusOK, e.Success, "affect")
	return

}
