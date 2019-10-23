package admin

import (
	"context"
	"fadmin/pkg/config"
	"github.com/xormplus/xorm"
)

type Dao struct {
	config.Config
	*xorm.Engine
}

func (d Dao) Delete2Table(beans [][2]interface{}) error {
	return d.delete2Table(beans)
}

func (d Dao) InsertMany(ctx context.Context, beans ...interface{}) error {
	return d.insertMany(beans...)
}

func (d Dao) JoinMany(ctx context.Context, bean interface{}, table string, query []string, values []interface{}, join [][3]interface{}) (int64, error) {
	return d.join2Table(bean, table, query, values, join)
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

func (d Dao) TxInsert(ctx context.Context, model interface{}) error {
	return d.insertOne(model)
	//switch model.(type) {
	//case admin.Domain:
	//	return d.insertOne(model)
	//}
	//return nil
}

func (d Dao) Exist(ctx context.Context, model interface{}) bool {
	return d.exist(model)
}

func (d Dao) Delete(ctx context.Context, id int64, model interface{}) (int64, error) {
	return d.delete(id, model)
	//switch model.(type) {
	//case admin.Domain:
	//	return d.delete(id, model)
	//}
	//return 0, nil

}
func NewDb(path string) *Dao {
	return &Dao{config.NewConfig(path), config.NewDb()}
}
