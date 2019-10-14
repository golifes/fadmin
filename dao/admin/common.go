package admin

import (
	"fadmin/tools/utils"
	"fmt"
	"log"
)

func (d Dao) query(sql string, fields []string, values []interface{}, model interface{}) (interface{}, error) {
	if rows, err := d.DB().Query(sql, values); err != nil {
		return nil, nil
	} else {
		cols := make([]interface{}, len(fields))
		for i, _ := range fields {
			cols[i] = &fields[i]
		}
		data := make([]map[string]interface{}, 0)
		if rows != nil {
			column, _ := rows.Columns()
			for rows.Next() {
				err = rows.Scan(cols...)
				row := make(map[string]interface{}) //每行数据
				for k, v := range fields {
					key := column[k]
					row[key] = v
				}
				data = append(data, row)
			}
			fmt.Println(data)
		}
		return data, nil
	}
}

func (d Dao) count(sql string) int {
	row := d.DB().QueryRow(sql)
	var count int64
	if err := row.Scan(&count); err != nil {
		return 0
	} else {
		return int(count)
	}
}
func (d Dao) insert(sql string, values []interface{}) int {
	stmt, err := d.DB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0
	}
	if result, err := stmt.Exec(values...); err != nil {
		return 0
	} else if rowsAffected, err := result.RowsAffected(); err != nil {
		return 0
	} else {
		return int(rowsAffected)
	}
}

func (d Dao) txInsert(values map[string][]interface{}) {
	tx, err := d.DB().Begin()
	if utils.CheckError(err, tx) {
		for k, v := range values {
			if _, err := tx.Exec(k, v); err != nil {
				log.Printf("%s", err)
				err = tx.Rollback()
				log.Printf("%s", err)
				return
			}
		}
		err = tx.Commit()
	}

}
