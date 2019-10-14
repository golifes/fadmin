package e

const (
	//通用编码
	Success          = 200
	Errors           = 500
	Forbid           = 403
	Unauthorized     = 401
	TokenCreateError = 2000
	JobError         = 3000

	NoLogin = 0001

	//业务编码
	AddError      = 1000 //添加失败
	UpdateError   = 1001 //更新失败
	DeleteError   = 1002 //删除失败
	FindError     = 1003 //查询失败
	EmptyError    = 1004 //查询为空
	ExistError    = 1005 //已经存在
	NotExistError = 1006 //用户不存在
	ParamError    = 1007 //参数错误
	ParamLose     = 1008 //参数丢失
)
