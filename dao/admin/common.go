package admin

import "fmt"

func (d Dao) query(sql string, fields []string, values []interface{}, model interface{}) (interface{}, int) {
	if rows, err := d.db.Query(sql, values); err != nil {
		return nil, 0
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
		return data, 0
	}
}

func (d Dao) count(sql string) int {
	row := d.db.QueryRow(sql)
	var count int64
	if err := row.Scan(&count); err != nil {
		return 0
	} else {
		return int(count)
	}
}
func (d Dao) insert(sql string, values []interface{}) int {
	stmt, err := d.db.Prepare(sql)
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
