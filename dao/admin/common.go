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
		_, err := tx.Session().Insert(bean)
		if err != nil {
			err = tx.RollbackTrans()
			return err
		}
	}
	err = tx.CommitTrans()
	return err
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
func (d Dao) findOne(bean interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64) {
	var count int64
	var err error
	if len(query) == 0 {
		count, err = d.Engine.Table(table).Limit(ps, ps*pn).OrderBy(orderBy).FindAndCount(bean)
		if utils.CheckError(err, count) {
			return bean, count
		}
		return nil, 0

	} else {
		count, err = d.Engine.Table(table).Where(strings.Join(query, " "), values...).Limit(ps, ps*pn).OrderBy(orderBy).FindAndCount(bean)
		if utils.CheckError(err, count) {
			return bean, count
		}
		return nil, 0
	}
}

func (d Dao) updateMap(table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	//affected, err := engine.Table(new(User)).Id(id).Update(map[string]interface{}{"age":0})
	return d.Engine.Table(table).Where(strings.Join(query, ""), values...).Cols(cols...).Update(m)
}
func (d Dao) updateStruct(bean interface{}, cols, query []string, values []interface{}) (int64, error) {
	return d.Engine.Where(strings.Join(query, ""), values...).Cols(cols...).Update(bean)
}

//2表联查
func (d Dao) join2Table(bean interface{}, table string, query []string, values []interface{}, join [][3]interface{}) (int64, error) {
	//mmm := [][3]interface{}{{
	//	"join", "type", "group.id = user.group_id"},
	//	{"join", []string{"group", "g"}, "group.id = user.group_id"},}

	session := d.Engine.Table(table)
	for _, v := range join {
		session.Join(v[0].(string), v[1], v[2].(string))
	}
	count, err := session.Where(strings.Join(query, " "), values...).FindAndCount(bean)
	//count, err := d.Engine.Table(table).Join(joinOperator, tableName, condition, args).Where(strings.Join(query, " "), values...).FindAndCount(bean)
	return count, err
}

//级联删除
func (d Dao) delete2Table(beans [][2]interface{}) error {
	/**
	数据结构如下
	[][2]{{id,model},{id,model}}
	*/
	session := d.Engine.NewSession()
	defer session.Close()
	tx, err := session.BeginTrans()
	if err != nil {
		return err
	}
	for _, bean := range beans {
		_, err := tx.Session().ID(bean[0]).Delete(&bean)
		if err != nil {
			tx.RollbackTrans()
		}
	}
	return nil
}
