package admin

type DbHandler interface {
	Counter
}

type Counter interface {
	Query(
		db []string, cols []string, query map[string]interface{}, pn, ps int, model interface{},
	) (interface{}, int) //表名,字段,条件,分页 ,返回值是结果集和message model是查询结果集的rows映射
	Count(db []string, query map[string]string) (int, int) //返回数据条数和message
}
type Handler interface {
	Insert(op string, db []string, rows map[string]interface{}) //自动添加创建时间  op表示操作，比如 select
	Update(op string, db []string, rows map[string]interface{}) //自动更新更新时间
	Delete(op string, db []string, rows map[string]interface{}) //删除操作
}

//var _ DbHandler = Dao(nil)

var _ DbHandler = Dao{}
