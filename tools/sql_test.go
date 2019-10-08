package tools

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	cols := []string{"w.a", "w.b"}
	db := []string{" wx w "}
	field := []string{"w.a =", "or w.b <=", "and w.c !="}
	queryValues := []interface{}{"1", 1, 10}
	fmt.Println(queryValues)
	sql := Select(db, cols, field)
	t.Log(sql)
}

func TestDelete(t *testing.T) {
	db := "wx"
	field := []string{"w.a =", "or w.b <=", "and w.c !="}
	sql := Delete(db, field)
	t.Log(sql)
}
func TestUpdate(t *testing.T) {
	db := "wx"
	field := []string{"w.a =", " w.b = "}
	query := []string{"w.a =", "or w.b <=", "and w.c !="}

	sql := Update(db, field, query)
	t.Log(sql)
}

func TestInsert(t *testing.T) {
	db := "wx"
	field := []string{"w.a =", " w.b = "}
	sql := Insert(db, field)
	t.Log(sql)
}
