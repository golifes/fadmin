package wx

import (
	"fadmin/model/wx"
	"fadmin/pkg/app"
	"fadmin/pkg/config"
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	"fmt"
	"net/http"
	"strings"
	"time"
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
	p.Forbid = 1
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

	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), wx.WeiXin{Key: p.Key, Uin: p.Uin}, []string{"key", "uin"}, []string{"biz = ? "}, []interface{}{p.Biz})
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
	query := []string{" forbid = ? "}
	values := []interface{}{1}
	if p.Name != "" {
		query = append(query, " and `name` like ? ")
		values = append(values, fmt.Sprintf("%s%s%s", "%", p.Name, "%"))
	}

	if p.Biz != "" {
		query = append(query, " and biz = ? ")
		values = append(values, p.Biz)
	}

	if p.Id != 0 {
		query = append(query, " and id = ? ")
		values = append(values, p.Id)
	}
	//type WeiXin struct {
	//	Id   int64  `json:"id"`
	//	Biz  string `json:"biz"`
	//	Name string `json:"name"`
	//}
	weiXin := make([]wx.RetWx, 0)

	list, count := h.logic.FindOne(g.NewContext(ctx), &weiXin, "wei_xin", "id desc ", query, values, ps, pn)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, e.Success, m)
}

func (h HttpWxHandler) ForBidWx(ctx app.GContext) {
	var p wx.ForBidWx
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), wx.WeiXin{Forbid: 1}, []string{"forbid", "mtime"}, []string{"id = ?"}, []interface{}{p.Id})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Id)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

func (h HttpWxHandler) WxList(ctx app.GContext) {
	var p wx.WxList
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	ps, pn := utils.Pagination(p.Ps, p.Pn, 10)
	query := []string{" forbid = ? "}
	values := []interface{}{1}

	if p.Title != "" {
		query = append(query, " and `title` like ? ")
		values = append(values, fmt.Sprintf("%s%s%s", "%", p.Title, "%"))
	}

	if p.Id != 0 {
		query = append(query, " and id = ? ")
		values = append(values, p.Id)
	}

	if p.StartTime.Unix() > 0 {
		query = append(query, " and public_time >= ? ")
		values = append(values, p.StartTime)
	}

	if p.EndTime.Unix() > 0 {
		query = append(query, " and public_time <= ? ")
		values = append(values, p.StartTime)
	}
	oderBy := "id desc "
	if p.OrderBy != "" {
		oderBy = p.OrderBy
	}
	weiXin := make([]wx.WeiXinList, 0)
	//
	list, count := h.logic.FindOne(g.NewContext(ctx), &weiXin, "wei_xin_list", oderBy, query, values, ps, pn)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, e.Success, m)
}

//获取多条数据
func (h HttpWxHandler) FindBizUinKey(ctx app.GContext) {
	//不接受任何参数
	var p wx.Ps
	g, err := h.common(ctx, &p)

	if err != nil {
		return
	}

	if p.Ps == 0 {
		p.Ps = 10
	}
	weiXin := make([]wx.WeiXinKey, 0)

	//这里需要改下，获取最近2小时未被抓取的公号
	list, count := h.logic.FindOne(g.NewContext(ctx), &weiXin, "wei_xin", "ctime desc ", nil, nil, p.Ps, 1)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, e.Success, m)

}

func (h HttpWxHandler) AddWxList(ctx app.GContext) {
	g := app.G{ctx}

	var p wx.ParamsAddWxList
	code := e.Success
	if !utils.CheckError(ctx.ShouldBindJSON(&p), "addList") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	ids := utils.FindBizStr(p.Url)
	biz := ""
	var w wx.WeiXinList
	if ids != nil {
		biz = ids[0]
		w.Id = p.Id
		w.Biz = biz
		w.ArticleId = utils.EncodeMd5(strings.Join(ids, "_"))
	} else {
		//log  url 提取idx等参数异常
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	//todo 先查询ArticleId是否存在，存在就不入库，不存在就存在入库
	if h.logic.Exist(g.NewContext(ctx), &w) {
		g.Json(http.StatusOK, e.Errors, "")
		return
	}
	w.Ptime = time.Unix(int64(p.Ptime), 0)
	w.Url = p.Url
	w.Title = p.Title
	if !utils.CheckError(h.logic.InsertOne(g.NewContext(ctx), w), "insert addList") {
		code = e.Errors
	}
	g.Json(http.StatusOK, code, "")
	return
}
