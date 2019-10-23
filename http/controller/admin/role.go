package adminc

import (
	"fadmin/model/admin"
	"fadmin/pkg/app"
	"fadmin/pkg/config"
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	"net/http"
)

/**
对角色的curd
*/

//添加角色
func (h HttpAdminHandler) AddRole(ctx app.GContext) {

	//先检查did  aid 是否存在 存在检查name是否存在
	var p admin.ParamsRole
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	exist := h.logic.Exist(g.NewContext(ctx), &admin.DomainApp{Did: p.Did, Id: p.Aid})
	if !exist {
		g.Json(http.StatusOK, e.DomainNotExist, p)
		return
	}
	rid := config.NewNodeId()
	err = h.logic.InsertMany(g.NewContext(ctx), &admin.DomainAppRole{Did: p.Did, Aid: p.Aid, Rid: rid, Status: 1}, &admin.Role{Id: rid, Name: p.Name, Status: 1, Rw: 1})
	if utils.CheckError(err, "addRole") {
		g.Json(http.StatusOK, e.Success, p.Name)
	} else {
		g.Json(http.StatusOK, e.Errors, p.Name)

	}

}
