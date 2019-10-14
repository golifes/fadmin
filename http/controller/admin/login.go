package adminc

import (
	"fadmin/model/admin"
	"fadmin/pkg/app"
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	"fmt"
	"net/http"
)

func (h HttpAdminHandler) Login(c app.GContext) {
	g := app.G{Context: c}

	var p admin.ParamsLogin
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "login") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	//先查询did aid是否存在,如果存在就查询这个用户是否存在
	domainApp := admin.DomainApp{Did: p.Did, Aid: p.Aid}
	values := []interface{}{p.Did, p.Aid}
	fields := []string{"did=", "aid="}

	fmt.Println(domainApp, fields, values)
	count, err := h.logic.Count(g.NewContext(c), "", fields, values)
	if !utils.CheckError(err, count) || count == 0 {
		g.Json(http.StatusOK, e.NotExistError, "")
	}

	//这里域和应用都存在，校验用户是否存在

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
