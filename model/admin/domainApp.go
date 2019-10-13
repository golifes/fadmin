package admin

import "time"

/**
Id 用Did和Aid计算出来一个值，然后代表当前唯一id,这样方便查询在插入数据的时候查询是否存在
*/

type DomainApp struct {
	Id     int64     `json:"id" `
	Name   string    `json:"name"  `
	Did    int64     `json:"did"  `
	Status int       `json:"status"  `
	Ctime  time.Time `json:"ctime" `
	Mtime  time.Time `json:"mtime" `
}
