package adminc

//func (h HttpAdminHandler) Login(ctx app.GContext) {
//
//	var p admin.ParamsLogin
//	g, err := h.common(ctx, &p)
//	if err != nil {
//		return
//	}
//	code := h.ExistDomainApp(g, ctx, p.Did, p.Aid, p)
//	if code != e.Success {
//		return
//	}
//
//	m := make(map[string]interface{})
//	m["token"] = ""
//	m["perm"] = []string{}
//	m["uid"] = 1
//	m["rid"] = 1
//	g.Json(http.StatusOK, e.Success, m)
//	return
//}
//
//func (h HttpAdminHandler) Register(ctx app.GContext) {
//	var p admin.ParamsLogin
//	g, err := h.common(ctx, &p)
//	if err != nil {
//		return
//	}
//
//	//校验是否存在
//	//values := []interface{}{p.Did, p.Aid}
//	//fields := []string{"did=", "aid="}
//	//
//	//count, err := h.logic.Count(g.NewContext(ctx), "", fields, values, p)
//	//if !utils.CheckError(err, count) || count == 0 {
//	//	g.Json(http.StatusOK, e.ParamError, "")
//	//	return
//	//}
//	code := h.ExistDomainApp(g, ctx, p.Did, p.Aid, p)
//	if code != e.Success {
//		return
//	}
//
//	//写数据到数据库
//	//fields := []string{"name", "pwd"}
//	//values := []interface{}{p.Name, utils.EncodeMd5(p.Pwd), p.Did, p.Aid}
//	//insert p 下面开始事物然后insert
//	//err = h.logic.TxInsert(g.NewContext(ctx), "", fields, values, p)
//	//if !utils.CheckError(err, "TxInsert") {
//	//	g.Json(http.StatusOK, e.Errors, err)
//	//	return
//	//}
//	g.Json(http.StatusOK, e.Success, p.Name)
//}
