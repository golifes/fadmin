package admin

import (
	"context"
	"fadmin/dao/admin"
)

type LogicHandler interface {
	Handler
}

type adminHandler interface {
}

type Handler interface {
	Query(ctx context.Context,
		table string, cols []string, fields []string, values []interface{}, pn, ps int, model interface{}) (interface{}, error) //表名,字段,条件,分页 ,返回值是结果集和message model是查询结果集的rows映射
	Count(ctx context.Context, db string, fields []string, values []interface{}, model interface{}) (int, error)                     //返回
	Insert(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) (int, error)                 //自动添加创建时间  op表示操作，比如 select
	Update(ctx context.Context, table string, query []string, fields []string, values []interface{}, model interface{}) (int, error) //自动更新更新时间
	Delete(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) (int, error)
}
type Logic struct {
	Db admin.DbHandler
}

func (l Logic) Query(ctx context.Context,
	table string, cols []string, fields []string, values []interface{}, pn, ps int, model interface{},
) (interface{}, error) {
	return l.Db.Query(ctx, table, cols, fields, values, pn, ps, model)
}

func (l Logic) Count(ctx context.Context, db string, fields []string, values []interface{}, model interface{}) (int, error) {
	return l.Db.Count(ctx, db, fields, values, model)
}

func (l Logic) Insert(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) (int, error) {
	return l.Db.Insert(ctx, table, fields, values, model)
}

func (l Logic) Update(ctx context.Context, table string, query []string, fields []string, values []interface{}, model interface{}) (int, error) {
	return l.Db.Update(ctx, table, query, fields, values, model)
}

func (l Logic) Delete(ctx context.Context, table string, fields []string, values []interface{}, model interface{}) (int, error) {
	return l.Db.Delete(ctx, table, fields, values, model)
}

var _ LogicHandler = Logic{}
