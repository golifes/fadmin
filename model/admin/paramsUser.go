package admin

/**
用户相关的参数model
*/

//login

type ParamsLogin struct {
	UserName string `json:"user_name"   binding:"required"`
	Pwd      string `json:"pwd"   binding:"required"`
	Did      string `json:"did" binding:"required"`
	Aid      string `json:"aid" binding:"required"`
}
