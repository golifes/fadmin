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

/**
删除
*/
func (h HttpAdminHandler) DeleteRole(ctx app.GContext) {
	var p admin.ParamsRoleDel
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	var role admin.DomainAppRole

	affect, err := h.logic.Delete(g.NewContext(ctx), p.Id, role)
	if err != nil {
		g.Json(http.StatusOK, e.DomainDeleteError, p.Id)
		return
	}
	g.Json(http.StatusOK, e.Success, affect)
}

/**
查询
*/

func (h HttpAdminHandler) FindRole(ctx app.GContext) {
	var p admin.ParamsRoleList
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	var query []string
	var values []interface{}
	if p.Rid != 0 {
		query, values = utils.Slice(query, values, " ar.rid = ? ", p.Rid)
	}
	if p.Did != 0 {
		query, values = utils.Slice(query, values, " ar.did = ? ", p.Did)
	}
	if p.Aid != 0 {
		query, values = utils.Slice(query, values, " ar.aid = ? ", p.Did)
	}
	if p.Name != "" {
		query, values = utils.Slice(query, values, " r.`name` like ? ", fmt.Sprintf("%s%s%s", "%", p.Name, "%"))
	}

	if p.Status != 0 {
		query, values = utils.Slice(query, values, " ar.status = ? ", p.Status)
	}

	ps, pn := utils.Pagination(p.Ps, p.Pn, 10)
	domain := make([]admin.ParamsRoleList, 0)
	//list, count := h.logic.FindOne(g.NewContext(ctx), &domain, table.Domain, " ar.ctime desc ", query, values, ps, pn)
	join := [][3]interface{}{{"inner", []string{"role", "r"}, "r.id = ar.rid"}}
	list, count := h.logic.FindMany(g.NewContext(ctx), &domain, table.DomainAppRole, "ar", "ar.id,ar.did,ar.aid,ar.rid,ar.status,ar.ctime,ar.mtime,r.name", "ar.ctime desc ", ps, pn, query, values, join)

	m := make(map[string]interface{})
	m["count"] = count
	m["data"] = list
	g.Json(http.StatusOK, e.Success, m)

}
