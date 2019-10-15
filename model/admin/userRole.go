package admin

import "time"

/**
user role
*/

type UserRole struct {
	Id     int64     `json:"id"  `
	Rid    int64     `json:"rid"  `
	Uid    int64     `json:"uid"  `
	Status int       `json:"status"  `
	Ctime  time.Time `json:"ctime" xorm:"created"` //创建时间
	Mtime  time.Time `json:"mtime" xorm:"updated"` //更新时间
}
