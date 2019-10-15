package admin

import "fadmin/tools/utils"

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

func (d Dao) exist(bean ...interface{}) error {
	exist, err := d.Engine.Exist(&bean)
	if exist && utils.CheckError(err, exist) {
		return nil
	}
	return err
}
