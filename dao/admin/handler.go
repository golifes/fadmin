package admin

import "context"

type DbHandler interface {
	Counter
	Handler
}

type Counter interface {
}
type Handler interface {
	TxInsert(ctx context.Context, model interface{}) error //自动添加创建时间  op表示操作，比如 select
	Exist(ctx context.Context, model interface{}) bool
	Delete(ctx context.Context, id int64, model interface{}) (int64, error)
	FindOne(ctx context.Context, model interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64)
	GetOne(ctx context.Context, model interface{}, cols ...string) interface{}
	UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error)
	UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error)
	FindMany(ctx context.Context, bean interface{}, table, alias, cols, orderBy string, ps, pn int, query []string, values []interface{}, join [][3]interface{}) (interface{}, int64)
	InsertMany(ctx context.Context, beans ...interface{}) error //事物
	DeleteMany(beans [][2]interface{}) error                    //事物
}

//var _ DbHandler = Dao(nil)

var _ DbHandler = Dao{}
