package admin

import "time"

type UserGroup struct {
	Gid    int64     `json:"gid"  `
	Uid    string    `json:"uid"  `
	Status int       `json:"status"  `
	Ctime  time.Time `json:"ctime"  ` //创建时间
	Mtime  time.Time `json:"mtime"  ` //更新时间
}
