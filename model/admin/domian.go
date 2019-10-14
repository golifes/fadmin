package admin

import "time"

/**
域
*/
type Domain struct {
	Id     int64     `json:"id"  `     //域id
	Name   string    `json:"name"  `   //名称
	Status int       `json:"status"  ` //状态 ,停用启用
	Ctime  time.Time `json:"ctime"  `  //创建时间
	Mtime  time.Time `json:"mtime"  `  //更新时间
	Did    string    `json:"did"`      //对外暴露的id
}
