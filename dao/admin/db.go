package admin

import (
	"database/sql"
	"fadmin/pkg/config"
)

type Dao struct {
	c  config.Config
	db *sql.DB
}

func NewDb(path string) *Dao {
	return &Dao{db: config.Db, c: config.NewConfig(path)}
}
