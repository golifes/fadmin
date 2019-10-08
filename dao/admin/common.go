package admin

import "fmt"

func (d Dao) query(sql string, values []interface{}) (interface{}, int) {
	result, err := d.db.Exec(sql)
	//d.db.Prepare()
	//d.db.PrepareContext()
	fmt.Println(result, err)
	return nil, 0
}
