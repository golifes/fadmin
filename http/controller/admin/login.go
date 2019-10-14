package adminc

import (
	"fadmin/model/admin"
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	"net/http"
)

func (h HttpAdminHandler) Login(ctx app.GContext) {

	var p admin.ParamsLogin
	g, err := h.common(ctx, p)
	if err != nil {
		return
	}
	//先查询did aid是否存在,如果存在就查询这个用户是否存在
	//domainApp := admin.DomainApp{Did: p.Did, Aid: p.Aid}
	values := []interface{}{p.Did, p.Aid}
	fields := []string{"did=", "aid="}
	//
	//count, err := h.logic.Count(g.NewContext(ctx), "", fields, values, p)
	//if !utils.CheckError(err, count) || count == 0 {
	//	g.Json(http.StatusOK, e.ParamError, "")
	//	return
	//}
	code := h.ExistDomainApp(g, ctx, p.Did, p.Aid, p)
	if code != e.Success {
		return
	}

	//这里域和应用都存在，校验用户是否存在
	cols := []string{"id"}
	fields = []string{"name", "pwd"}
	values = []interface{}{p.Name, utils.EncodeMd5(p.Pwd)}
	query, err := h.logic.Query(g.NewContext(ctx), "", cols, fields, values, 0, 0, p)
	if !utils.CheckError(err, query) {
		g.Json(http.StatusOK, e.UserNotExist, "")
		return
	}
	//用户存在,返回对应的信息(这里包括权限等信息)

	m := make(map[string]interface{})
	m["token"] = ""
	m["perm"] = []string{}
	m["uid"] = 1
	m["rid"] = 1
	g.Json(http.StatusOK, e.Success, m)
	return

	//先检查是否存在
	//domain := admin.Domain{Id: utils.EncodeMd5(p.Name)}

	//if !h.logic.Exist(&domain, nil) {
	//	domain.Name = p.Name
	//	domain.Status = 1
	//	code = h.logic.Add(domain)
	//
	//} else {
	//	code = e.ExistError
	//}
	//g.Json(http.StatusOK, code, "")
	//return
}

func (h HttpAdminHandler) Register(ctx app.GContext) {
	//g := app.G{Context: ctx}

	var p admin.ParamsLogin
	//code := e.Success
	//if !utils.CheckError(ctx.ShouldBindJSON(&p), "login") {
	//	code = e.ParamError
	//	g.Json(http.StatusOK, code, "")
	//	return
	//}
	g, err := h.common(ctx, p)
	if err != nil {
		return
	}
	//校验是否存在
	//values := []interface{}{p.Did, p.Aid}
	//fields := []string{"did=", "aid="}
	//
	//count, err := h.logic.Count(g.NewContext(ctx), "", fields, values, p)
	//if !utils.CheckError(err, count) || count == 0 {
	//	g.Json(http.StatusOK, e.ParamError, "")
	//	return
	//}
	code := h.ExistDomainApp(g, ctx, p.Did, p.Aid, p)
	if code != e.Success {
		return
	}

	//写数据到数据库
	fields := []string{"name", "pwd"}
	values := []interface{}{p.Name, utils.EncodeMd5(p.Pwd), p.Did, p.Aid}
	//insert p 下面开始事物然后insert
	affect, err := h.logic.Insert(g.NewContext(ctx), "", fields, values, p)
	if utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.Errors, err)
		return
	}
	g.Json(http.StatusOK, e.Success, "")
}
