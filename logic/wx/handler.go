package wx

import (
	"context"
	"fadmin/dao/wx"
)

type LogicHandler interface {
	handler
}

type handler interface {
	InsertOne(ctx context.Context, model interface{}) error
	Exist(ctx context.Context, model interface{}) bool
	UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error)
	UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error)
	FindOne(ctx context.Context, model interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64)
}

type Logic struct {
	Db wx.WeiXinHandler
}

func (l Logic) FindOne(ctx context.Context, model interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64) {
	return l.Db.FindOne(ctx, model, table, orderBy, query, values, ps, pn)
}

func (l Logic) UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error) {
	return l.Db.UpdateStruct(ctx, model, cols, query, values)
}

func (l Logic) UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	return l.Db.UpdateMap(ctx, table, m, cols, query, values)
}

func (l Logic) Exist(ctx context.Context, model interface{}) bool {
	return l.Db.Exist(ctx, model)
}

func (l Logic) InsertOne(ctx context.Context, model interface{}) error {
	return l.Db.InsertOne(ctx, model)
}

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: wx.NewDb(path)}
}
