package admin

/**
用户组
*/
type Group struct {
	Id        int64  `json:"id"  `        //用户组id
	GroupName string `json:"groupName"  ` //用户组名称
	Did       string `json:"did"  `       //和域关联，为了查询方便
}
