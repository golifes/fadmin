package admin

import "time"

/**
domain app user关联关系表(admin用)
*/

type DomainAppUser struct {
	Id     int64     `json:"id" `      //主键id
	Did    string    `json:"did" `     //域id
	Aid    string    `json:"aid" `     //应用id
	Uid    string    `json:"uid" `     //用户id
	Status int       `json:"status"  ` //状态
	CTime  time.Time `json:"ctime"  `  //创建时间
	MTime  time.Time `json:"mtime"  `  //更新时间
}
