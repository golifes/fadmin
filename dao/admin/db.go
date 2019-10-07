package admin

import (
	"database/sql"
	"fadmin/pkg/config"
	"fadmin/tools"
)

type Dao struct {
	c  config.Config
	db *sql.DB
}

func (d Dao) Query(db []string, cols []string, query map[string]interface{}, pn, ps int, model interface{}) (interface{}, int) {
	exec := tools.Exec("select", db, cols, query)
	d.query(exec)
	return nil, 0
}

func (d Dao) Count(db []string, query map[string]string) (int, int) {
	panic("implement me")
}

func NewDb(path string) *Dao {
	return &Dao{db: config.Db, c: config.NewConfig(path)}
}
