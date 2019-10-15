package admin

import (
	"context"
	"fadmin/model/admin"
	"fadmin/pkg/config"
	"github.com/xormplus/xorm"
)

type Dao struct {
	config.Config
	*xorm.Engine
}

func (d Dao) TxInsert(ctx context.Context, model interface{}) error {
	switch model.(type) {
	case admin.Domain:
		return d.insertOne(model)
	}
	return nil
}

func (d Dao) Exist(ctx context.Context, model interface{}) error {
	switch model.(type) {
	case admin.Domain:
		return d.exist(model)
	}
	return nil
}

func NewDb(path string) *Dao {
	return &Dao{config.NewConfig(path), config.NewDb()}
}
