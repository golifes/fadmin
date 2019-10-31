package wx

import "time"

type WeiXinList struct {
	Id        int64     `json:"id"`                      //公号id
	Biz       string    `json:"biz"`                     //
	ArticleId string    `json:"article_id"`              //文章id
	Title     string    `json:"title"`                   //标题
	Digest    string    `json:"digest"`                  //摘要
	Url       string    `json:"url"`                     //url
	SourceUrl string    `json:"source_url"`              //source_url
	Forbid    int       `json:"forbid" xorm:"default 1"` //是否被禁用
	Ptime     time.Time `json:"ptim"`                    //发布时间
	Ctime     time.Time `json:"ctime" xorm:"created"`    //创建时间
	Mtime     time.Time `json:"mtime" xorm:"updated"`    //更新时间
	Ps        int       `json:"ps" xorm:"-"`             //
	Pn        int       `json:"pn" xorm:"-"`             //
}

//func (w WeiXinList) Deadline() (deadline time.Time, ok bool) {
//	panic("implement me")
//}
//
//func (w WeiXinList) Done() <-chan struct{} {
//	panic("implement me")
//}
//
//func (w WeiXinList) Err() error {
//	panic("implement me")
//}
//
//func (w WeiXinList) Value(key interface{}) interface{} {
//	panic("implement me")
//}
//
