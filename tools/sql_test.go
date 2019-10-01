package tools

import (
	"testing"
)

func TestExec(t *testing.T) {
	cols := []string{"w.a", "w.b"}
	m := map[string]string{" from  ": " wx w "}
	sql := Exec(cols, m)
	t.Log(sql)
}
