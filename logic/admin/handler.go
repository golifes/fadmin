package admin

import (
	"context"
	"fadmin/dao/admin"
)

type LogicHandler interface {
	Handler
}

type Handler interface {
	TxInsert(ctx context.Context, model interface{}) error
	Exist(ctx context.Context, model interface{}) bool
	Delete(ctx context.Context, id int64, model interface{}) (int64, error)
	FindOne(ctx context.Context, model interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64)
	GetOne(ctx context.Context, model interface{}, cols ...string) interface{}

	UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error)
	UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error)
	FindMany(ctx context.Context, bean interface{}, table, alias, cols, orderBy string, ps, pn int, query []string, values []interface{}, join [][3]interface{}) (interface{}, int64)
	InsertMany(ctx context.Context, beans ...interface{}) error
	DeleteMany(beans [][2]interface{}) error
	//	table string, cols []string, fields []string, values []interface{}, pn, ps int, model interface{}, orderBy string) (interface{}, error) //表名,字段,条件,分页 ,返回值是结果集和message model是查询结果集的rows映射
	//Count(ctx context.Context, db string, fields []string, values []interface{}, model interface{}) (int, error)                                //返回
	//TxInsert(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) error                                 //自动添加创建时间  op表示操作，比如 select
	//Update(ctx context.Context, table string, query []string, fields []string, values []interface{}, model interface{}) (int, error)            //自动更新更新时间
	//Delete(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) (int, error)
}
type Logic struct {
	Db admin.DbHandler
}

func (l Logic) GetOne(ctx context.Context, model interface{}, cols ...string) interface{} {
	return l.Db.GetOne(ctx, model, cols...)
}

func (l Logic) DeleteMany(beans [][2]interface{}) error {
	return l.Db.DeleteMany(beans)
}

func (l Logic) InsertMany(ctx context.Context, beans ...interface{}) error {
	return l.Db.InsertMany(ctx, beans...)
}

func (l Logic) FindMany(ctx context.Context, bean interface{}, table, alias, cols, orderBy string, ps, pn int, query []string, values []interface{}, join [][3]interface{}) (interface{}, int64) {
	return l.Db.FindMany(ctx, bean, table, alias, cols, orderBy, ps, pn, query, values, join)
}

func (l Logic) UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error) {
	return l.Db.UpdateStruct(ctx, model, cols, query, values)
}

func (l Logic) UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	return l.Db.UpdateMap(ctx, table, m, cols, query, values)
}

func (l Logic) FindOne(ctx context.Context, model interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64) {
	return l.Db.FindOne(ctx, model, table, orderBy, query, values, ps, pn)
}

func (l Logic) TxInsert(ctx context.Context, model interface{}) error {
	return l.Db.TxInsert(ctx, model)
}
func (l Logic) Exist(ctx context.Context, model interface{}) bool {
	return l.Db.Exist(ctx, model)
}
func (l Logic) Delete(ctx context.Context, id int64, model interface{}) (int64, error) {
	return l.Db.Delete(ctx, id, model)
}

//func (l Logic) Query(ctx context.Context,
//	table string, cols []string, fields []string, values []interface{}, pn, ps int, model interface{}, orderBy string) (interface{}, error) {
//	return l.Db.Query(ctx, table, cols, fields, values, pn, ps, model, orderBy)
//}
//
//func (l Logic) Count(ctx context.Context, db string, fields []string, values []interface{}, model interface{}) (int, error) {
//	return l.Db.Count(ctx, db, fields, values, model)
//}
//
//func (l Logic) TxInsert(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) error {
//	return l.Db.TxInsert(ctx, table, fields, values, model)
//}
//
//func (l Logic) Update(ctx context.Context, table string, query []string, fields []string, values []interface{}, model interface{}) (int, error) {
//	return l.Db.Update(ctx, table, query, fields, values, model)
//}
//
//func (l Logic) Delete(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) (int, error) {
//	return l.Db.Delete(ctx, table, fields, values, model)
//}

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: admin.NewDb(path)}
}
