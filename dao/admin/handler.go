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
	FineOne(ctx context.Context, ps, pn int, query []string, values []interface{}, bean interface{}) (interface{}, int64)
	UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error)
	UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error)
}

//var _ DbHandler = Dao(nil)

var _ DbHandler = Dao{}
