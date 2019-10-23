package admin

/**
domain app user关联关系表(admin用)
*/

type DomainAppUser struct {
	Id     int64 `json:"id" `                  //主键id
	Did    int64 `json:"did" `                 //域id
	Aid    int64 `json:"aid" `                 //应用id
	Uid    int64 `json:"uid" `                 //用户id
	Status int   `json:"status"  `             //状态
	Ctime  int   `json:"ctime" xorm:"created"` //创建时间
	Mtime  int   `json:"mtime" xorm:"updated"` //更新时间
}
