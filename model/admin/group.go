package admin

/**
用户组
*/
type Group struct {
	Id        int64  `json:"id"  `                 //用户组id
	GroupName string `json:"groupName"  `          //用户组名称
	Did       int64  `json:"did"  `                //和域关联，为了查询方便
	Status    int    `json:"status"  `             //状态 ,停用启用
	Ctime     int    `json:"ctime" xorm:"created"` //创建时间
	Mtime     int    `json:"mtime" xorm:"updated"` //更新时间
}
