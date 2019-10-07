package admin

import "fmt"

func (d Dao) query(sql string) {
	result, err := d.db.Exec(sql)
	d.db.Prepare()
	d.db.PrepareContext()
	fmt.Println(result, err)
}
