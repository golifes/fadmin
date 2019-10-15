package admin

import "time"

/**
role  perms
*/

type RolePerms struct {
	Id     int64     `json:"id"  `
	Rid    int64     `json:"rid" `
	Pid    int64     `json:"pid"  `
	Status int       `json:"status"  `
	Ctime  time.Time `json:"ctime" xorm:"created"` //创建时间
	Mtime  time.Time `json:"mtime" xorm:"updated"` //更新时间
}
