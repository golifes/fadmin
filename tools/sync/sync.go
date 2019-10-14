package main

import (
	"fadmin/model/admin"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

func main() {

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"abc123456",
		"127.0.0.1",
		"fadmin")
	fmt.Println(dns)
	EngDb, err := xorm.NewEngine("mysql", dns)

	ping := EngDb.Ping()
	if ping != nil || err != nil {
		panic(ping)
	}

	EngDb.Sync2(new(admin.DomainApp))
	EngDb.Sync2(new(admin.DomainAppRole))
	EngDb.Sync2(new(admin.DomainAppUser))
	EngDb.Sync2(new(admin.Domain))
	EngDb.Sync2(new(admin.Group))
	EngDb.Sync2(new(admin.Role))
	EngDb.Sync2(new(admin.User))
	EngDb.Sync2(new(admin.UserRole))
	EngDb.Sync2(new(admin.UserGroup))

}
