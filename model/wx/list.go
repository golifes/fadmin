package wx

import "time"

type WeiXinList struct {
	ArticleId  string    `json:"article_id"`           //文章id
	Title      string    `json:"title"`                //标题
	Digest     string    `json:"digest"`               //摘要
	Url        string    `json:"url"`                  //url
	SourceUrl  string    `json:"source_url"`           //sourceurl
	PublicTime time.Time `json:"public_time"`          //发布时间
	Ctime      time.Time `json:"ctime" xorm:"created"` //创建时间
	Mtime      time.Time `json:"mtime" xorm:"updated"` //更新时间
}
