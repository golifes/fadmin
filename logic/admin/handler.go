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
		db []string, cols []string, query map[string]interface{}, pn, ps int, model interface{},
	) (interface{}, int) //表名,字段,条件,分页 ,返回值是结果集和message model是查询结果集的rows映射
	Insert(ctx context.Context, op string, db []string, rows map[string]interface{}) //自动添加创建时间  op表示操作，比如 select
	Update(ctx context.Context, op string, db []string, rows map[string]interface{}) //自动更新更新时间
	Delete(ctx context.Context, op string, db []string, rows map[string]interface{}) //删除操作
}

type Logic struct {
	Db admin.DbHandler
}

func (l Logic) Query(ctx context.Context,
	db []string, cols []string, query map[string]interface{}, pn, ps int, model interface{},
) (interface{}, int) {
	return l.Db.Query(ctx, db, cols, query, pn, ps, model)
}

func (l Logic) Insert(ctx context.Context, op string, db []string, rows map[string]interface{}) {
	panic("implement me")
}

func (l Logic) Update(ctx context.Context, op string, db []string, rows map[string]interface{}) {
	panic("implement me")
}

func (l Logic) Delete(ctx context.Context, op string, db []string, rows map[string]interface{}) {
	panic("implement me")
}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: admin.NewDb(path)}
}

var _ LogicHandler = Logic{}
