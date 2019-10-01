package admin

import (
	"database/sql"
	"fadmin/pkg/config"
)

type Dao struct {
	c  config.Config
	db *sql.DB
}

func (d Dao) Query(db map[string]string, cols []string, query map[string]interface{}, pn, ps int) {
	panic("implement me")
}

func (d Dao) Count(db map[string]string, query map[string]string) {
	panic("implement me")
}

func NewDb(path string) *Dao {
	return &Dao{db: config.Db, c: config.NewConfig(path)}
}
