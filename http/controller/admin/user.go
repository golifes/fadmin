package adminc

import (
	"fadmin/model/admin"
	"fadmin/pkg/app"
	"fadmin/pkg/config"
	"fadmin/pkg/e"
	"fadmin/pkg/jwt"
	"fadmin/tools/utils"
	"net/http"
)

func (h HttpAdminHandler) Register(ctx app.GContext) {
	var p admin.ParamsLogin
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

	uid := config.NewNodeId()
	err = h.logic.InsertMany(g.NewContext(ctx), &admin.DomainAppUser{Did: p.Did, Aid: p.Aid, Uid: uid, Status: 1}, &admin.User{Id: uid, Name: p.Name, Pwd: utils.EncodeMd5(p.Pwd), Did: p.Did, Aid: p.Aid})
	if utils.CheckError(err, "register") {
		g.Json(http.StatusOK, e.Success, p.Name)
	} else {
		g.Json(http.StatusOK, e.Errors, p.Name)

	}
}

func (h HttpAdminHandler) Login(ctx app.GContext) {
	var p admin.ParamsLogin
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

	user := h.logic.GetOne(g.NewContext(ctx), &admin.User{Name: p.Name, Pwd: utils.EncodeMd5(p.Pwd), Did: p.Did, Aid: p.Aid, Status: 1}, "id", "name")
	if user == nil {
		g.Json(http.StatusOK, e.UserNotExist, p.Name)
		return

	}

	u := user.(*admin.User)
	//uid := u.Id
	//fmt.Println(uid)
	token, err := jwt.GenerateToken(u.Name, u.Id, u.Did, u.Aid, 1)
	m := make(map[string]interface{})
	m["token"] = token
	m["uid"] = u.Id
	m["did"] = u.Did
	m["aid"] = u.Aid
	if utils.CheckError(err, token) {
		g.Json(http.StatusOK, e.Success, m)
		return
	}
	g.Json(http.StatusOK, e.CreateTokenError, p.Name)

}
