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

func (d Dao) FineOne(ctx context.Context, ps, pn int, query []string, values []interface{}, bean interface{}) (interface{}, int64) {
	switch bean.(type) {
	case admin.ParamsDomainList:
		var ret []admin.Domain
		//bean interface{}, cond ...interface{}
		return d.find(&ret, ps, pn, query, values, bean)
	}
	return nil, 0
}

func (d Dao) TxInsert(ctx context.Context, model interface{}) error {
	switch model.(type) {
	case admin.Domain:
		return d.insertOne(model)
	}
	return nil
}

func (d Dao) Exist(ctx context.Context, model interface{}) bool {
	switch model.(type) {
	case *admin.Domain:
		return d.exist(model)
	}
	return false
}

func (d Dao) Delete(ctx context.Context, id int64, model interface{}) (int64, error) {
	switch model.(type) {
	case admin.Domain:
		return d.delete(id, model)
	}
	return 0, nil

}
func NewDb(path string) *Dao {
	return &Dao{config.NewConfig(path), config.NewDb()}
}
