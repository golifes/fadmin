package admin

type UserGroup struct {
	Uid    int64 `json:"uid" binding:"required"`
	Gid    int64 `json:"gid" binding:"required"`
	Status int   `json:"status"`               //
	Ctime  int   `json:"ctime" xorm:"created"` //创建时间
	Mtime  int   `json:"mtime" xorm:"updated"` //更新时间
}
