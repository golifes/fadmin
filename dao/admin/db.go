package admin

import (
	"context"
	"fadmin/model/admin"
	"fadmin/pkg/config"
	"fadmin/tools/utils"
	"fmt"
	"github.com/golifes/sqlo"
)

type Dao struct {
	config.Config
	sqlo.Engine
}

func (d Dao) Query(ctx context.Context,
	table string, col []string, fields []string, values []interface{}, pn, ps int, model interface{}) (interface{}, error) {
	sql := d.Select(col...).From(table).OrderBy(" id desc").Limit(ps, pn).String()
	switch model.(type) {
	case admin.User:
		return d.query(sql, fields, values, admin.User{})

	}
	return nil, nil
}

func (d Dao) Count(ctx context.Context, db string, fields []string, values []interface{}, model interface{}) (int, error) {
	switch model.(type) {
	case admin.ParamsLogin:
		sql := d.Select("name").From(db).Where("did=?").And("aid=?").String()
		fmt.Println(sql)
		return d.count(sql, values...), nil
	}
	return 0, nil
}

func (d Dao) TxInsert(ctx context.Context, db string, fields []string, values []interface{}, model interface{}) error {

	switch model.(type) {
	case admin.ParamsLogin:
		sqlList := map[string][]interface{}{}
		userSql := d.Insert("user").Cols("name", "pwd", "uid").AddNowTime([]string{"mtime", "ctime"}).String()
		userValue := values[:2]
		uid := utils.EncodeMd5(utils.StringJoin(values[0].(string), "_", values[1].(string)))
		userValue = append(userValue, uid)
		sqlList[userSql] = userValue
		domainAppUser := d.Insert("domain_app_user").Cols("did", "aid", "uid", "status").AddNowTime([]string{"mtime", "ctime"}).String()
		domainAppValue := values[2:]
		domainAppValue = append(domainAppValue, uid, 1)
		sqlList[domainAppUser] = domainAppValue
		return d.txInsert(sqlList)
	}

	return nil
}

func (d Dao) Update(ctx context.Context, db string, query []string, fields []string, values []interface{}, model interface{}) (int, error) {
	panic("implement me")
}

func (d Dao) Delete(ctx context.Context, db string, fields []string, values []interface{}, model interface{}) (int, error) {
	panic("implement me")
}

func NewDb(path string) *Dao {
	return &Dao{config.NewConfig(path), config.NewDb()}
}
