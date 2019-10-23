package admin

/**
role  perms
*/

type RolePerms struct {
	Id     int64 `json:"id"  `
	Rid    int64 `json:"rid" `
	Pid    int64 `json:"pid"  `
	Status int   `json:"status"  `
	Ctime  int   `json:"ctime" xorm:"created"` //创建时间
	Mtime  int   `json:"mtime" xorm:"updated"` //更新时间
}
