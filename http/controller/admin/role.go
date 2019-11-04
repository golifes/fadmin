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
角色和用户绑定,和部门没关系，部门只是一个归档作用
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
	join := [][3]interface{}{{"inner", []string{"role", "r"}, "r.id = ar.rid"}, {"inner", []string{"domain_app", "a"}, "a.id = ar.aid"}}
	list, count := h.logic.FindMany(g.NewContext(ctx), &domain, table.DomainAppRole, "ar", "ar.id,ar.did,a.name as aname,ar.rid,ar.status,ar.ctime,ar.mtime,r.name", "ar.ctime desc ", ps, pn, query, values, join)

	m := make(map[string]interface{})
	m["count"] = count
	m["data"] = list
	g.Json(http.StatusOK, e.Success, m)

}

/**
更新角色：只能修改这个角色属于那个应用
*/

func (h HttpAdminHandler) UpdateRole(ctx app.GContext) {
	var p admin.ParamsRoleUpdate
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	/**
	先查询did 和aid是否是属于同一个域下面的
	*/
	exist := h.logic.Exist(g.NewContext(ctx), &admin.DomainAppRole{Did: p.Did, Id: p.Aid})
	if !exist {
		g.Json(http.StatusOK, e.DomainNotExist, "")
		return

	}
	//拼接更新字段
	var cols []string
	if p.Status != 0 {
		cols = append(cols, "status")
	}
	if p.Aid != 0 {
		cols = append(cols, "aid")
	}

	var role admin.DomainAppRole
	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), role, cols, []string{" id = ? "}, []interface{}{p.Id})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Id)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

/**
修改角色名称
*/

func (h HttpAdminHandler) UpdateRoleName(ctx app.GContext) {
	var p admin.ParamsRoleName
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	exist := h.logic.Exist(g.NewContext(ctx), &admin.Role{Id: p.Rid})
	if !exist {
		g.Json(http.StatusOK, e.DomainNotExist, "")
		return

	}

	var role admin.DomainAppRole
	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), role, []string{"name"}, []string{" id = ? "}, []interface{}{p.Rid})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.Errors, p.Name)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

/**
给人绑定角色 userRole
params:	{

	"uid":"",#用户id
	"rid":"",#角色id
}
*/

func (h HttpAdminHandler) AddUserRole(ctx app.GContext) {

	/**
	todo
	aid did在用户信息中获取
	*/
	var p admin.ParamsUserRole
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	exist := h.logic.Exist(g.NewContext(ctx), &admin.DomainAppRole{Rid: p.Rid})
	if !exist {
		g.Json(http.StatusOK, e.DomainNotExist, "")
		return
	}
	id := config.NewNodeId()
	err = h.logic.InsertMany(g.NewContext(ctx), &admin.UserRole{
		Id:     id,
		Rid:    p.Rid,
		Uid:    p.Uid,
		Status: 1,
	})
}
