package router

func (e *Engine) weiXin() {
	r := e.Group("/wx")
	{

		r.POST("/add", e.AddWx)       //添加公众号
		r.POST("/key", e.UpdateWxKey) //更新公众号key
		r.GET("/biz", e.FindWxBiz)    //获取公号列表数据
		r.POST("/forbid", e.ForBidWx) //禁用微信
		r.GET("/list", e.WxList)
		r.GET("/key", e.FindBizUinKey)

		//差一个入库详情页和列表数据接口
	}
}
