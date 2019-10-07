package tools

import (
	"testing"
)

func TestExec(t *testing.T) {
	cols := []string{"w.a", "w.b"}
	db := []string{" wx w "}
	queryMap := make(map[string]interface{})
	queryMap[" w.a = "] = 1
	queryMap[" w.b <= "] = 2
	queryMap[" w.c != "] = "aaa"
	sql := Exec("delete ", db, cols, queryMap)
	t.Log(sql)
}
