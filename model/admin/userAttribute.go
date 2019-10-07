package admin

/**
用户属性,abac
*/

type UserAttr struct {
	Uid  int64  `json:"uid"  `
	Attr string `json:"arrt"` //数据库存json数据
}
