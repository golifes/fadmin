package admin

type ParamsDomainList struct {
	Id     int64 `json:"id"`
	Ps     int
	Pn     int    //分页
	Name   string //模糊查询
	Status int
}

type ParamsId struct {
	Id int64 `json:"id"  binding:"required"`
}

type ParamsIdName struct {
	Id     int64  `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Status int    `json:"status"`
}

type ParamsAppList struct {
	Id     int64  `json:"id"`
	Did    int64  `json:"did"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Ps     int    `json:"ps"`
	Pn     int    `json:"pn"`
}
