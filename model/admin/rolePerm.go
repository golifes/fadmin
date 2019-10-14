package admin

import "time"

/**
role  perms
*/

type RolePerms struct {
	Id     int64     `json:"id"  `
	Rid    string    `json:"rid" `
	Pid    string    `json:"pid"  `
	Status int       `json:"status"  `
	Ctime  time.Time `json:"ctime"  ` //创建时间
	Mtime  time.Time `json:"mtime"  ` //更新时间`
}
