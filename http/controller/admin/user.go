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

/**
手机号码登录
*/

func (h HttpAdminHandler) LoginPhone(ctx app.GContext) {

	/**
	查询手机号码是否存在,存在就发验证码
	*/
	var p admin.ParamsPhoneLogin
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

	if p.Code != "1234" {
		g.Json(http.StatusOK, e.CodeError, p.Code)
		return
	}

	user := h.logic.GetOne(g.NewContext(ctx), &admin.User{Phone: p.Phone, Did: p.Did, Aid: p.Aid}, "id")

	m := make(map[string]interface{})
	var uid int64
	if user == nil {
		uid = config.NewNodeId()
		err = h.logic.InsertMany(g.NewContext(ctx), &admin.User{Id: uid, Phone: p.Phone, Pwd: utils.EncodeMd5(p.Phone), Did: p.Did, Aid: p.Aid})
		if !utils.CheckError(err, "insert") {
			g.Json(http.StatusOK, e.Errors, p.Phone)
			return
		}
	} else {
		u := user.(*admin.User)
		if u.Status != 1 {
			g.Json(http.StatusOK, e.Forbid, p.Phone)
			return
		}
		uid = u.Id
	}

	token, err := jwt.GenerateToken(p.Phone, uid, p.Did, p.Aid, 1)
	if !utils.CheckError(err, token) {
		g.Json(http.StatusOK, e.CreateTokenError, p.Phone)
		return
	}
	m["token"] = token
	m["uid"] = uid
	m["did"] = p.Did
	m["aid"] = p.Aid
	g.Json(http.StatusOK, e.Success, m)
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

//禁用用户
func (h HttpAdminHandler) ForbidUser(ctx app.GContext) {
	var p admin.ParamsIds
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), admin.User{Status: p.Status}, []string{"status"}, []string{" id = ? "}, []interface{}{p.Id})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Id)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

//更新密码
func (h HttpAdminHandler) UpdatePwd(ctx app.GContext) {
	var p admin.ParamsPwdUpdate
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), admin.User{Pwd: utils.EncodeMd5(p.Pwd)}, []string{"pwd"}, []string{" id = ? "}, []interface{}{p.Id})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Id)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

//更新手机号码
func (h HttpAdminHandler) UpdatePhone(ctx app.GContext) {

	var p admin.ParamsPhoneUpdate
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), admin.User{Phone: p.Phone}, []string{"phone"}, []string{" id = ? "}, []interface{}{p.Id})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Id)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

func (h HttpAdminHandler) DeleteUser(ctx app.GContext) {
	var p admin.ParamsId

	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	if p.Id == 0 {
		g.Json(http.StatusOK, e.ParamError, "")
		return
	}
	var user admin.User
	affect, err := h.logic.Delete(g.NewContext(ctx), p.Id, user)
	if err != nil {
		g.Json(http.StatusOK, e.DomainDeleteError, p.Id)
		return
	}
	g.Json(http.StatusOK, e.Success, affect)
}
