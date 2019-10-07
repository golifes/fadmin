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
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"   `
}
