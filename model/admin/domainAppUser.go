package admin

import "time"

/**
domain app user关联关系表(admin用)
*/

type DomainAppUser struct {
	Id     int64     `json:"id"  `
	Did    string    `json:"did"  `
	Aid    string    `json:"aid"  `
	Uid    string    `json:"uid"  `
	Status int       `json:"status"  `
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"  `
}
