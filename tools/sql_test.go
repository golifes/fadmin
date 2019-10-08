package tools

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestSelect(t *testing.T) {
	cols := []string{"w.a", "w.b"}
	db := []string{" wx w "}
	field := []string{"w.a =", "or w.b <=", "and w.c !="}
	queryValues := []interface{}{"1", 1, 10}
	fmt.Println(queryValues)
	sql := Select(db, cols, field, 0, 10, "id desc")
	t.Log(sql)
}

func TestDelete(t *testing.T) {
	db := "wx"
	field := []string{"w.a =", "or w.b <=", "and w.c !="}
	sql := Delete(db, field)
	t.Log(sql)
}
func TestUpdate(t *testing.T) {
	//db := "wx"
	//field := []string{"w.a =", " w.b = "}
	//query := []string{"w.a =", "or w.b <=", "and w.c !="}
	//
	//sql := Update(db, field, query)
	//t.Log(sql)
	var i interface{}
	var buf bytes.Buffer
	i = 10
	buf.WriteString(string(i.(int)))
	fmt.Println(1, buf.String())
	fmt.Println(1, buf.Bytes())

}

func TestInsert(t *testing.T) {
	db, err := sql.Open("mysql", "root:abc123456@tcp(58.87.64.219)/fadmin?charset=utf8&parseTime=True&loc=Local") //连接数据库
	fmt.Println(err)
	rows, err := db.Query("select id ,name  from  domain_app  limit  10; ")
	fields := []string{"id", "name"}
	values := make([]interface{}, len(fields))
	for i, _ := range fields {
		values[i] = &fields[i]
	}
	data := make([]map[string]interface{}, 0)
	if rows != nil {
		column, _ := rows.Columns()
		for rows.Next() {
			err = rows.Scan(values...)
			row := make(map[string]interface{}) //每行数据
			for k, v := range fields {
				key := column[k]
				row[key] = v
			}
			data = append(data, row)
		}
		fmt.Println(data)
	}
	var count int64
	row := db.QueryRow("select count(1) from domain_app ;")
	row.Scan(&count)
	fmt.Println(count)

}
