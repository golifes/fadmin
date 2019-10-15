package admin

import "time"

/**
域
*/
type Domain struct {
	Id     int64     `json:"id"  `                       //域id
	Name   string    `json:"name"  binding:"required"  ` //名称
	Status int       `json:"status"  `                   //状态 ,停用启用
	Ctime  time.Time `json:"ctime" xorm:"created"`       //创建时间
	Mtime  time.Time `json:"mtime" xorm:"updated"`       //更新时间
}
