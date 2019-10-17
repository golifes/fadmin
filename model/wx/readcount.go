package wx

import "time"

type WeiXinCount struct {
	Id         int64     //我们维护的公号id
	ArticleId  string    //文章id
	ReadCount  int       //阅读量
	ThumbCount int       //点赞量
	Ctime      time.Time `json:"ctime" xorm:"created"` //创建时间
	Mtime      time.Time `json:"mtime" xorm:"updated"` //更新时间
}
