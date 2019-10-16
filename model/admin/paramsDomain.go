package admin

type ParamsDomainList struct {
	Id   int64 `json:"id"`
	Ps   int
	Pn   int    //分页
	Name string //模糊查询
}
