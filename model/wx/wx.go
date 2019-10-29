package wx

import "time"

/**
微信公号
*/
type WeiXin struct {
	Id         int64     `json:"id"`                       //主键Id
	WxId       string    `json:"wx_id" binding:"required"` //微信id
	Name       string    `json:"name" binding:"required"`  //微信名称
	Url        string    `json:"url" binding:"required"`   //微信头像
	Desc       string    `json:"desc" binding:"required"`  //公号描述
	Biz        string    `json:"biz" binding:"required"`   //公号biz
	Count      string    `json:"count"`                    //公号文章数量
	Forbid     int       `json:"forbid" xorm:"default 1"`  //是否被禁用
	Key        string    `json:"key"`                      //公号Key
	Uin        string    `json:"uin"`                      //用户唯一标识
	Ctime      time.Time `json:"ctime" xorm:"created"`     //创建时间
	Mtime      time.Time `json:"mtime" xorm:"updated"`     //更新时间
	SpiderTime time.Time `json:"spider_time"`              //最后一次抓取时间
	Note       string    `json:"note"`                     //备用字段
}

type WxApi struct {
	Id   int64  `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

//阅读历史

type Ua struct {
	Id int64  `json:"id"`
	Ua string `json:"ua"`
}
