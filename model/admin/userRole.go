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
	Ctime  time.Time `json:"ctime"  ` //创建时间
	Mtime  time.Time `json:"mtime"  ` //更新时间
}
