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

	exist = h.logic.Exist(g.NewContext(ctx), &admin.Group{Did: p.Did, Aid: p.Aid, Name: p.Name})
	if exist {
		g.Json(http.StatusOK, e.GroupExist, "")
		return
	}
	p.Id = config.NewNodeId()
	err = h.logic.InsertMany(g.NewContext(ctx), p)
	if err != nil {
		g.Json(http.StatusOK, e.Errors, p.Name)
	} else {
		g.Json(http.StatusOK, e.Success, p.Name)
	}
}

//删除组
func (h HttpAdminHandler) DeleteGroup(ctx app.GContext) {
	var p admin.ParamsId

	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	affect, err := h.logic.Delete(g.NewContext(ctx), p.Id, admin.Group{Id: p.Id})
	if err != nil {
		g.Json(http.StatusOK, e.Errors, p.Id)
		return
	}
	g.Json(http.StatusOK, e.Success, affect)
}

//更新组  修改组名称
func (h HttpAdminHandler) UpdateGroup(ctx app.GContext) {
	var p admin.ParamsIdName
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	exist := h.logic.Exist(g.NewContext(ctx), &admin.Group{Name: p.Name})
	if exist {
		g.Json(http.StatusOK, e.DomainExist, "")
		return
	}
	cols := []string{"name"}
	m := make(map[string]interface{})
	m["name"] = p.Name

	if p.Status != 0 {
		cols = append(cols, "status")
		m["status"] = p.Status
	}
	affect, err := h.logic.UpdateMap(g.NewContext(ctx), table.Group, m, cols, []string{"id = ? "}, []interface{}{p.Id})
	//affect, err := h.logic.UpdateStruct(g.NewContext(ctx), admin.Group{Name:p.Name}, cols, []string{"id = ? "}, []interface{}{p.Id})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Id)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

//查询组信息
func (h HttpAdminHandler) FindGroup(ctx app.GContext) {
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
	domain := make([]admin.Group, 0)
	list, count := h.logic.FindOne(g.NewContext(ctx), &domain, table.Group, " ctime desc ", query, values, ps, pn)
	m := make(map[string]interface{})
	m["count"] = count
	m["data"] = list
	g.Json(http.StatusOK, e.Success, m)
}
