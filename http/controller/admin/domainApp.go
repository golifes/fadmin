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
查询域下面的app信息
*/

func (h HttpAdminHandler) AddApp(ctx app.GContext) {
	var p admin.DomainApp
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	//先判断did是否存在
	exist := h.logic.Exist(g.NewContext(ctx), &admin.Domain{Id: p.Did})
	if exist {
		exist = h.logic.Exist(g.NewContext(ctx), &admin.DomainApp{Name: p.Name})
		if exist {
			g.Json(http.StatusOK, e.AppExist, "")
			return
		}
	} else {
		g.Json(http.StatusOK, e.DomainNotExist, "")
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

//删除
func (h HttpAdminHandler) DeleteApp(ctx app.GContext) {
	var p admin.ParamsId

	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	if p.Id == 0 {
		g.Json(http.StatusOK, e.ParamError, "")
		return
	}
	var domainApp admin.DomainApp
	affect, err := h.logic.Delete(g.NewContext(ctx), p.Id, domainApp)
	if err != nil {
		g.Json(http.StatusOK, e.DomainDeleteError, p.Id)
		return
	}
	g.Json(http.StatusOK, e.Success, affect)
	return
}

/**
修改 不允许修改这个应用的did
*/

func (h HttpAdminHandler) UpdateApp(ctx app.GContext) {
	var p admin.DomainApp
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	if p.Id == 0 {
		g.Json(http.StatusOK, e.ParamError, "")
		return
	}
	exist := h.logic.Exist(g.NewContext(ctx), &admin.DomainApp{Name: p.Name})
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

func (h HttpAdminHandler) FindApp(ctx app.GContext) {

}
