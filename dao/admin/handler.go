package admin

import "context"

type DbHandler interface {
	Counter
	Handler
}

type Counter interface {
	Query(ctx context.Context,
		table string, cols []string, fields []string, values []interface{}, pn, ps int, model interface{},
	) (interface{}, error) //表名,字段,条件,分页 ,返回值是结果集和message model是查询结果集的rows映射
	Count(ctx context.Context, db string, fields []string, values []interface{}, model interface{}) (int, error) //返回数据条数和message
}
type Handler interface {
	InsertTable(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) (int, error)            //自动添加创建时间  op表示操作，比如 select
	Update(ctx context.Context, table string, query []string, fields []string, values []interface{}, model interface{}) (int, error) //自动更新更新时间
	Delete(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) (int, error)                 //删除操作
}

//var _ DbHandler = Dao(nil)

var _ DbHandler = Dao{}
