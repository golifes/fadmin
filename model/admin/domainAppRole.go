package admin

import "time"

/**
domain app role关联关系(admin用)
*/

type DomainAppRole struct {
	Id     int64     `json:"id" `     //主键id
	Did    string    `json:"did" `    //域id
	Aid    string    `json:"aid" `    //应用id
	Rid    string    `json:"rid" `    //角色id
	Status int       `json:"status" ` //状态 停用启用
	Ctime  time.Time `json:"ctime"  ` //创建时间
	Mtime  time.Time `json:"mtime"  ` //更新时间
}
