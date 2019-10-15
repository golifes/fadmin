package admin

import "time"

type UserGroup struct {
	Gid    int64     `json:"gid"  `
	Uid    int64     `json:"uid"  `
	Status int       `json:"status"  `
	Ctime  time.Time `json:"ctime" xorm:"created"` //创建时间
	Mtime  time.Time `json:"mtime" xorm:"updated"` //更新时间
}
