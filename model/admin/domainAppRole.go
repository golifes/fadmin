package admin

/**
domain app role关联关系(admin用)
*/

type DomainAppRole struct {
	Id     int64 `json:"id" `                  //主键id
	Did    int64 `json:"did" `                 //域id
	Aid    int64 `json:"aid" `                 //应用id
	Rid    int64 `json:"rid" `                 //角色id
	Status int   `json:"status" `              //状态 停用启用
	Ctime  int   `json:"ctime" xorm:"created"` //创建时间
	Mtime  int   `json:"mtime" xorm:"updated"` //更新时间
}
