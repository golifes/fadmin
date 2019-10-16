package tools

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	db, _ := NewDb("root:abc123456@tcp(58.87.64.219)/fadmin?charset=utf8&parseTime=True&loc=Local")
	i := db.
		Select("a", "b").From("wx").
		Where().
		And("name", "age").
		string()
	fmt.Println(i)

}
