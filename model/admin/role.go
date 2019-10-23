package admin

/**
role

角色分可读 可写 读写
*/

type Role struct {
	Id     int64  `json:"id"  `
	Name   string `json:"name" `
	Status int    `json:"status"  `
	Rw     int    `json:"rw"  `    //读写标志  1为读 2为写 3可读可写
	Ctime  int    `json:"ctime"  ` //创建时间
	Mtime  int    `json:"mtime"  ` //更新时间
}
