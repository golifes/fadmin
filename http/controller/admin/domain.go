package adminc

import (
	"fadmin/model/admin"
	"fadmin/pkg/app"
	"fadmin/pkg/config"
	"fadmin/pkg/e"
	"fadmin/pkg/table"
	"fadmin/tools/utils"
	"fmt"
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

/**
参数{"id":1,"name":2}
*/
func (h HttpAdminHandler) DeleteDomain(ctx app.GContext) {
	var p admin.Domain

	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	if p.Id == 0 {
		g.Json(http.StatusOK, e.ParamError, "")
		return
	}
	affect, err := h.logic.Delete(g.NewContext(ctx), p.Id, p)
	if err != nil {
		g.Json(http.StatusOK, e.DomainDeleteError, p.Id)
		return
	}
	g.Json(http.StatusOK, e.Success, affect)
}

func (h HttpAdminHandler) UpdateDomain(ctx app.GContext) {
	var p admin.Domain
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	if p.Id == 0 {
		g.Json(http.StatusOK, e.ParamError, "")
		return
	}
	exist := h.logic.Exist(g.NewContext(ctx), &admin.Domain{Name: p.Name})
	if exist {
		g.Json(http.StatusOK, e.DomainExist, "")
		return
	}
	cols := []string{"name"}
	if p.Status != 0 {
		cols = append(cols, "status")
	}

	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), p, cols, []string{"id = ? "}, []interface{}{p.Id})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Id)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

//查询还的优化
func (h HttpAdminHandler) FindDomain(ctx app.GContext) {
	var p admin.ParamsDomainList
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	var query []string
	var values []interface{}
	if p.Id != 0 {
		query, values = utils.Slice(query, values, " id = ? ", p.Id)
	}

	if p.Name != "" {
		query, values = utils.Slice(query, values, " `name` like ? ", fmt.Sprintf("%s%s%s", "%", p.Name, "%"))
	}

	if p.Status != 0 {
		query, values = utils.Slice(query, values, " status = ? ", p.Status)
	}

	ps, pn := utils.Pagination(p.Ps, p.Pn, 10)
	domain := make([]admin.Domain, 0)
	list, count := h.logic.FindOne(g.NewContext(ctx), &domain, table.Domain, " ctime desc ", query, values, ps, pn)
	m := make(map[string]interface{})
	m["count"] = count
	m["data"] = list
	g.Json(http.StatusOK, e.Success, m)

}
