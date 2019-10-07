package admin

import "time"

/**
perms
设计做成一级目录 二级目录 三级目录
	api
api 分 读 写
这个是针对角色的
*/

type Menu struct {
	Id         int64     `json:"id"  `        //主键id
	Category   int       `json:"category"`    //类型  1为目录 2为api
	Name       string    `json:"name"`        //名称
	Icon       string    `json:"icon"`        //目录的icon
	ParentId   string    `json:"parent_id"`   //父id
	ParentPath string    `json:"parent_path"` //父路径
	Sort       int       `json:"sort"`        //子菜单排序
	Hidden     int       `json:"hidden"`      //是否展示
	Status     int       `json:"status" `
	CTime      time.Time `json:"ctime" `
	MTime      time.Time `json:"mtime"  `
	Action     Actions   `json:"action"`
	Source     Sources   `json:"source"`
	Did        string    `json:"did"  ` //冗余字段，为了查询方便
	Aid        string    `json:"adi"  `
}

type Actions []*Action
type Sources []*Source

//菜单动作对象 id和上面的id是同一个id
type Action struct {
	Id    int64     `json:"id"`
	Mid   string    `json:"id"  `   //权限id
	Code  string    `json:"code"  ` //动作编号
	Name  string    `json:"name"`   //名称
	CTime time.Time `json:"ctime"  `
	MTime time.Time `json:"mtime" `
}

//菜单资源对象
type Source struct {
	Id     int64     `json:"id"`
	Mid    string    `json:"pid"  `  //权限id
	Code   string    `json:"code"  ` //资源编号
	Name   string    `json:"name"`   //名称
	Method string    `json:"method"` //请求方式
	Uri    string    `json:"uri"`    //请求uri
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"  `
}

/**
结构体对象转换
*/

//func (m Menu)ToMenu()*Menu  {
//	item := &Menu{
//		Id:m.Id,
//	}
//}
