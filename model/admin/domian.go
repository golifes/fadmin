package admin

import "time"

/**
域
*/
type Domain struct {
	Id     int64     `json:"id"  `
	Name   string    `json:"name"  `
	Status int       `json:"status"  `
	Ctime  time.Time `json:"ctime"  `
	Mtime  time.Time `json:"mtime"  `
}
