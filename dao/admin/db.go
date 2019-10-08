package admin

import (
	"context"
	"database/sql"
	"fadmin/pkg/config"
	"fadmin/tools"
)

type Dao struct {
	c  config.Config
	db *sql.DB
}

func (d Dao) Query(ctx context.Context,
	db []string, cols []string, fields []string, values []interface{}, pn, ps int, model interface{},
) (interface{}, int) {
	s := tools.Select(db, cols, fields)
	return d.query(s, values)
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
	return &Dao{db: config.Db, c: config.NewConfig(path)}
}
