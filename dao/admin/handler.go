package admin

type DbHandler interface {
	Counter
}

type Counter interface {
	Query(db map[string]string, cols []string, query map[string]interface{}, pn, ps int) //表名,字段,条件,分页
	Count(db map[string]string, query map[string]string)
}

var _ DbHandler = Dao(nil)
