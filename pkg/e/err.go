package e

var msgDict = map[int]string{
	Success:           "成功",
	Errors:            "异常",
	Forbid:            "您暂无权限",
	Unauthorized:      "认证失败，请重新登录",
	EmptyError:        "查询结果为空",
	ParamError:        "参数错误",
	ParamLose:         "缺少必要的参数",
	NoLogin:           "请登录后查看更多页",
	DomainExist:       "域已经存在",
	DomainDeleteError: "域删除失败",
}

func GetMsg(code int) string {
	if msg, ok := msgDict[code]; ok {
		return msg
	}
	return msgDict[Errors]
}
