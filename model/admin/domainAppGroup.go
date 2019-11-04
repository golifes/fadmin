package admin

/**
用户组 给部门一个默认角色
*/
type Group struct {
	Id     int64  `json:"id"  `                     //用户组id
	Name   string `json:"name"  binding:"required"` //用户组名称
	Did    int64  `json:"did"  binding:"required"`  //和域关联，为了查询方便
	Aid    int64  `json:"aid"  binding:"required"`
	Rid    int64  `json:"rid" binding:"required"`
	Status int    `json:"status"  `             //状态 ,停用启用
	Ctime  int    `json:"ctime" xorm:"created"` //创建时间
	Mtime  int    `json:"mtime" xorm:"updated"` //更新时间
}
