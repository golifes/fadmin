package e

var msgDict = map[int]string{
	Success:          "成功",
	Errors:           "异常",
	Forbid:           "您暂无权限",
	Unauthorized:     "认证失败，请重新登录",
	TokenCreateError: "token异常",
	JobError:         "注册任务失败",
	AddError:         "添加失败",
	UpdateError:      "更新失败",
	DeleteError:      "删除失败",
	FindError:        "查询失败",
	EmptyError:       "查询结果为空",
	ExistError:       "已存在",
	NotExistError:    "不存在",
	ParamError:       "参数错误",
	ParamLose:        "缺少必要的参数",
	NoLogin:          "请登录后查看更多页",
}

func GetMsg(code int) string {
	if msg, ok := msgDict[code]; ok {
		return msg
	}
	return msgDict[Errors]
}
