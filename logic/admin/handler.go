package admin

import "fadmin/dao/admin"

type LogicHandler interface {
	adminHandler
}

type adminHandler interface {
}

type Logic struct {
	Db admin.DbHandler
}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: admin.NewDb(path)}
}
