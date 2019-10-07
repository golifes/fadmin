package tools

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func Exec(op string, cols []string, db []string, queryMap map[string]interface{}) string {
	var buf bytes.Buffer

	buf.WriteString(op)
	buf.WriteString(strings.Join(cols, ","))
	buf.WriteString(" from ")
	buf.WriteString(strings.Join(db, ","))

	if queryMap != nil {
		buf.WriteString(" where ")
	}
	count := 0
	for k, v := range queryMap {
		buf.WriteString(k)
		switch v.(type) {
		case int, int64:
			buf.WriteString(fmt.Sprintf("%d", v))
		case string:
			buf.WriteString(v.(string))
		case reflect.MapIter:
			fmt.Println("not support")
		}
		if len(queryMap)-1 != count {
			buf.WriteString(" and ")
		}
		count += 1

	}
	return buf.String()
}
