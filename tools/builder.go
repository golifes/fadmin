package tools

import (
	"bytes"
	"database/sql"
)

type tableInfo struct {
	Db     string   //表明
	Fields []string //查询字段
	Cols   []string //条件字段
	Values []string //条件值
	S      string
}

type DB struct {
	engine *sql.DB
	s      string
	err    error
}

func NewDb(dns string) (db DB, err error) {
	db = DB{}
	open, err := sql.Open("mysql", dns)
	db.engine = open //连接数据库
	db.err = err
	return
}

func (d DB) Select(cols ...string) DB {
	var buf bytes.Buffer
	buf.WriteString(" select ")
	count := 0
	for _, v := range cols {
		if len(cols)-1 != count {
			buf.WriteString(v)
			buf.WriteString(",")
		} else {
			buf.WriteString(v)
		}
	}
	d.s = buf.String()
	return d
}

func (d DB) From(db string) DB {
	var buf bytes.Buffer
	buf.WriteString(d.s)

	buf.WriteString(" from ")
	buf.WriteString(db)

	d.s = buf.String()
	return d
}

func (d DB) Where() DB {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	buf.WriteString(" where ")
	d.s = buf.String()
	return d
}
func (d DB) And(fields ...string) DB {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	count := 0
	for _, v := range fields {
		if len(fields)-1 != count {
			buf.WriteString(v)
			buf.WriteString(" = ?  and ")
		} else {
			buf.WriteString(v)
			buf.WriteString(" = ?  ")
		}
		count += 1
	}
	d.s = buf.String()
	return d
}

func (d DB) string() string {
	return d.s
}
