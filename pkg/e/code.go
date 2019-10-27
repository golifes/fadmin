package e

const (
	//通用编码
	Success      = 200
	Errors       = 500
	Forbid       = 403 //被禁用
	Unauthorized = 401 //未授权

	CreateTokenError = 1000
	CodeError        = 1010
	NoLogin          = 10000 //未登录
	RequestError     = 10001
	ParamLose        = 10010
	ParamError       = 10020
	//user
	UserNotExist     = 100000 //用户不存在
	UserExist        = 100001 //用户已存在
	RegisterError    = 100002 //注册异常
	UpdatePwdError   = 100003 //修改密码异常
	UpdatePhoneError = 100004 //修改手机号码异常
	UpdateUserError  = 100005 //修改用户信息异常
	DeleteUserError  = 100010 //删除用户异常
	EmptyError       = 100020 //查询结果集异常

	DomainNotExist    = 100030 //
	DomainExist       = 100031
	DomainDeleteError = 100032
	AppExist          = 100040
	WxExist           = 100100
	UpdateWxError     = 100101

	GroupExist = 100110
)
