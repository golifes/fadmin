package admin

import (
	"context"
	"fadmin/model/admin"
	"fadmin/pkg/config"
	"github.com/golifes/sqlo"
)

type Dao struct {
	config.Config
	*sqlo.Engine
}

func (d Dao) Query(ctx context.Context,
	table string, col []string, fields []string, values []interface{}, pn, ps int, model interface{},
) (interface{}, int) {
	sql := d.Select(col...).From(table).OrderBy(" id desc").Limit(ps, pn).String()
	switch model.(type) {
	case admin.User:
		return d.query(sql, fields, values, admin.User{})

	}
	return nil, 0
}

func (d Dao) Count(ctx context.Context, db []string, fields []string, values []interface{}) (int, int) {
	panic("implement me")
}

func (d Dao) Insert(ctx context.Context, db string, fields []string, values []interface{}) {
	panic("implement me")
}

func (d Dao) Update(ctx context.Context, db string, query []string, fields []string, values []interface{}) {
	panic("implement me")
}

func (d Dao) Delete(ctx context.Context, db string, fields []string, values []interface{}) {
	panic("implement me")
}

func NewDb(path string) *Dao {
	return &Dao{config.NewConfig(path), config.NewDb()}
}
