package tools

import (
	"bytes"
	"strings"
)

func Select(db []string, cols []string, fields []string) string {
	var buf bytes.Buffer
	buf.WriteString(" select ")
	buf.WriteString(strings.Join(cols, ","))
	buf.WriteString(" from ")
	buf.WriteString(strings.Join(db, ","))
	if fields != nil {
		buf.WriteString(" where ")
	}
	for _, v := range fields {
		buf.WriteString(v)
		buf.WriteString(" ? ")
	}
	return buf.String()
}
func Delete(db string, fields []string) string {
	var buf bytes.Buffer
	buf.WriteString(" delete from ")
	buf.WriteString(db)

	if fields != nil {
		buf.WriteString(" where ")
	}
	for _, v := range fields {
		buf.WriteString(v)
		buf.WriteString(" ? ")
	}
	return buf.String()
}

//UPDATE table_name SET field1=new-value1, field2=new-value2
func Update(db string, fields []string, query []string) string {
	var buf bytes.Buffer
	buf.WriteString(" update ")
	buf.WriteString(db)
	buf.WriteString(" set ")
	for _, v := range fields {
		buf.WriteString(v)
		buf.WriteString(" ? ")
	}
	buf.WriteString(" where ")
	for _, v := range query {
		buf.WriteString(v)
		buf.WriteString(" ? ")
	}
	return buf.String()
}

func Insert(db string, fields []string) string {
	var buf bytes.Buffer
	buf.WriteString(" insert into ")
	buf.WriteString(db)
	buf.WriteString(" ( ")
	buf.WriteString(strings.Join(fields, ","))
	buf.WriteString(" ) values ( ")
	count := 0
	for index, _ := range fields {
		if index != count-1 {
			buf.WriteString(" ? ")
		} else {
			buf.WriteString(" ? ")
			buf.WriteString(",")
		}
	}
	buf.WriteString(" ) ")

	return buf.String()
}
