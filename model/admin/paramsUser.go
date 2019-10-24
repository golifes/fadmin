package admin

/**
用户相关的参数model
*/

//login

type ParamsLogin struct {
	Name string `json:"name"   binding:"required"`
	Pwd  string `json:"pwd"   binding:"required"`
	Did  int64  `json:"did" binding:"required"`
	Aid  int64  `json:"aid" binding:"required"`
}

type ParamsPwdUpdate struct {
	Id  int64  `json:"id"  binding:"required"`
	Pwd string `json:"pwd"  binding:"required"`
}

type ParamsPhoneUpdate struct {
	Id    int64  `json:"id"  binding:"required"`
	Phone string `json:"phone"  binding:"required"`
}

type ParamsPhoneLogin struct {
	Phone string `json:"phone"  binding:"required"`
	Code  string `json:"code" binding:"required"`
	Did   int64  `json:"did" binding:"required"`
	Aid   int64  `json:"aid" binding:"required"`
}
