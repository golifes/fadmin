package admin

import "time"

/**
user role
*/

type UserRole struct {
	Id     int64     `json:"id"  `
	Rid    string    `json:"rid"  `
	Uid    string    `json:"uid"  `
	Status int       `json:"status"  `
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"  `
}
