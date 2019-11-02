package wx

import "time"

type WeiXinKey struct {
	Id   int64  `json:"id"`
	WxId string `json:"wx_id"`
	Biz  string `json:"biz" binding:"required"`
	Key  string `json:"key" binding:"required"`
	Uin  string `json:"uin" binding:"required"`
}

type Wx struct {
	Id   int64  `json:"id"`   //通过id查询
	Name string `json:"name"` //通过name模糊查询
	Biz  string `json:"biz"`  //通过biz查询
	Ps   int
	Pn   int
}

type RetWx struct {
	Id   int64  `json:"id" binding:"required"`
	Biz  string `json:"biz" `
	Name string `json:"name"`
}

type ForBidWx struct {
	Id string `json:"id" binding:"required"`
}

//获取微信文章列表数据
type WxList struct {
	Id        int64     `json:"id"`                      //公号id
	ArticleId string    `json:"article_id"`              //文章id
	Title     string    `json:"title"`                   //标题
	Forbid    int       `json:"forbid" xorm:"default 1"` //是否被禁用
	StartTime time.Time `json:"start_time"`              //发布时间
	EndTime   time.Time `json:"end_time"`                //结束时间
	Ps        int       `json:"ps"`                      //分页
	Pn        int       `json:"pn"`                      //分页
	OrderBy   string    `json:"order_by"`                //排序
}

type Ps struct {
	Ps int `json:"ps"`
}

type ParamsAddWxList struct {
	Id    int64  `json:"id"`
	Url   string `json:"url"  binding:"required" `   //文章url
	Title string `json:"title"  binding:"required" ` //文章标题
	Ptime int64  `json:"ptime"`                      //发布时间
}
