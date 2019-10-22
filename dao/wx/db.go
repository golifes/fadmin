package wx

import (
	"context"
	"errors"
	"fadmin/model/wx"
	"fadmin/pkg/config"
	"github.com/xormplus/xorm"
)

type Dao struct {
	config.Config
	*xorm.Engine
}

func NewDb(path string) *Dao {
	return &Dao{config.NewConfig(path), config.NewDb()}
}

func (d Dao) InsertOne(ctx context.Context, model interface{}) error {
	switch model.(type) {
	case wx.WeiXin:
		return d.insertOne(model)
	}

	return errors.New("不支持的操作")
}

func (d Dao) Exist(ctx context.Context, model interface{}) bool {
	switch model.(type) {
	case *wx.WeiXin:
		return d.exist(model)
	}

	return true
}

func (d Dao) UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error) {
	return d.updateStruct(model, cols, query, values)
}

func (d Dao) UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	return d.updateMap(table, m, cols, query, values)
}

func (d Dao) FindOne(ctx context.Context, model interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64) {
	return d.findOne(model, table, orderBy, query, values, ps, pn)
}
