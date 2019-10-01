package tools

import (
	"bytes"
	"strings"
)

func Exec(cols []string, db map[string]string) string {
	var buf bytes.Buffer

	buf.WriteString("select ")
	buf.WriteString(strings.Join(cols, ","))

	for k, v := range db {
		buf.WriteString(k)
		buf.WriteString(v)
	}

	return buf.String()

}
