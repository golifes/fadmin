package wx

import "context"

type WeiXinHandler interface {
	handler
}

type handler interface {
	InsertOne(ctx context.Context, model interface{}) error
	Exist(ctx context.Context, model interface{}) bool
	UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error)
	UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error)
	FindOne(ctx context.Context, model interface{}, ps, pn int, orderBy, table string, query []string, values []interface{}) (interface{}, int64)
}

var _ WeiXinHandler = Dao{}
