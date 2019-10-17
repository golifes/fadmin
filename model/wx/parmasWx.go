package wx

type WeiXinKey struct {
	Biz string `json:"biz" binding:"required"`
	Key string `json:"key" binding:"required"`
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
