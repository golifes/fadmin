package wx

import (
	"fadmin/model/wx"
	"fadmin/pkg/app"
	"fadmin/pkg/config"
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	"fmt"
	"net/http"
)

/**
页面添加微信号
params:{"wx_id":"","name":"","url":"","desc":"","biz":""}
*/
func (h HttpWxHandler) AddWx(ctx app.GContext) {
	var p wx.WeiXin
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	exist := h.logic.Exist(g.NewContext(ctx), &wx.WeiXin{Biz: p.Biz})
	if exist {
		g.Json(http.StatusOK, e.WxExist, p.Name)
		return
	}
	p.Id = config.NewNodeId()
	err = h.logic.InsertOne(g.NewContext(ctx), p)

	if err != nil {
		g.Json(http.StatusOK, e.Errors, p.Name)
	} else {
		g.Json(http.StatusOK, e.Success, p.Name)
	}
}

func (h HttpWxHandler) UpdateWxKey(ctx app.GContext) {
	var p wx.WeiXinKey
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), wx.WeiXin{Key: p.Key}, []string{"key"}, []string{"biz = ? "}, []interface{}{p.Biz})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Biz)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

func (h HttpWxHandler) FindWxBiz(ctx app.GContext) {
	var p wx.Wx
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	ps, pn := utils.Pagination(p.Ps, p.Pn, 10)
	var query []string
	var values []interface{}
	if p.Name != "" {
		query = append(query, " `name` like ? ")
		values = append(values, fmt.Sprintf("%s%s%s", "%", p.Name, "%"))
	}

	if p.Biz != "" {
		if len(values) == 0 {
			query = append(query, "  biz = ? ")
		} else {
			query = append(query, " and biz = ? ")
		}
		values = append(values, p.Biz)
	}

	if p.Id != 0 {
		if len(values) == 0 {
			query = append(query, "  id = ? ")
		} else {
			query = append(query, " and id = ? ")
		}
		values = append(values, p.Id)
	}
	//type WeiXin struct {
	//	Id   int64  `json:"id"`
	//	Biz  string `json:"biz"`
	//	Name string `json:"name"`
	//}
	weiXin := make([]wx.RetWx, 0)

	list, count := h.logic.FindOne(g.NewContext(ctx), &weiXin, ps, pn, "id desc ", "wei_xin", query, values)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, e.Success, m)

}
