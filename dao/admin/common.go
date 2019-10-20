package admin

import (
	"fadmin/tools/utils"
	"strings"
)

func (d Dao) insertOne(beans ...interface{}) error {
	_, err := d.Engine.Insert(beans)
	return err
}

func (d Dao) insertMany(beans ...interface{}) error {
	session := d.Engine.NewSession()
	defer session.Close()
	tx, err := session.BeginTrans()
	if err != nil {
		return err
	}
	for _, bean := range beans {

		_, err := tx.Session().Insert(&bean)
		if err != nil {
			tx.RollbackTrans()
		}
	}
	return nil
}

func (d Dao) exist(bean ...interface{}) bool {
	//exist, err := d.Engine.Exist(&admin.Domain{Name:"用户管理"})
	exist, err := d.Engine.Exist(bean...)

	if exist && utils.CheckError(err, exist) {
		return true
	}
	return false
}
func (d Dao) delete(id int64, bean interface{}) (int64, error) {
	return d.Engine.Id(id).Delete(bean)
}

//要支持模糊查询
func (d Dao) find(bean interface{}, ps, pn int, query []string, values []interface{}, cond ...interface{}) (interface{}, int64) {
	var count int64
	var err error
	if len(query) == 0 {
		count, err = d.Engine.FindAndCount(bean, cond...)
		if err != nil {
			return nil, 0
		}
	} else {
		//d.Engine.Where()
	}

	return bean, count
}

func (d Dao) updateMap(table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	//affected, err := engine.Table(new(User)).Id(id).Update(map[string]interface{}{"age":0})
	return d.Engine.Table(table).Where(strings.Join(query, ""), values...).Cols(cols...).Update(m)
}
func (d Dao) updateStruct(bean interface{}, cols, query []string, values []interface{}) (int64, error) {
	return d.Engine.Where(strings.Join(query, ""), values...).Cols(cols...).Update(bean)
}
