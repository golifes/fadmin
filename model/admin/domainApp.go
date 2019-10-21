package admin

/**
Id 用Did和Aid计算出来一个值，然后代表当前唯一id,这样方便查询在插入数据的时候查询是否存在
*/

type DomainApp struct {
	Id     int64  `json:"id" `                      //主键id
	Name   string `json:"name" binding:"required" ` //名称id
	Did    int64  `json:"did" binding:"required" `  //域id(用的对外暴露的id)
	Status int    `json:"status"  `                 //状态 停用启用
	Ctime  int    `json:"ctime" xorm:"created"`     //创建时间
	Mtime  int    `json:"mtime" xorm:"updated"`     //更新时间
}
