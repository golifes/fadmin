package admin

/**
user role
*/

type UserRole struct {
	Id     int64 `json:"id"  `
	Rid    int64 `json:"rid"  `
	Uid    int64 `json:"uid"  `
	Status int   `json:"status"  `
	Ctime  int   `json:"ctime" xorm:"created"` //创建时间
	Mtime  int   `json:"mtime" xorm:"updated"` //更新时间
}
