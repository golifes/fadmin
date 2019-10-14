package admin

import "time"

type UserGroup struct {
	Gid    int64     `json:"gid"  `
	Uid    string    `json:"uid"  `
	Status int       `json:"status"  `
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"  `
}
