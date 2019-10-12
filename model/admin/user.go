package admin

import "time"

/**
加上冗余两个字段
*/
type User struct {
	//Id  int64
	Id    int64     `json:"id"  `    //主键id
	Name  string    `json:"name"  `  //用户名
	Nick  string    `json:"name"  `  //昵称
	Pwd   string    `json:"pwd" `    //密码
	Phone string    `json:"phone"`   //电话号码
	CTime time.Time `json:"ctime"  ` //创建时间
	MTime time.Time `json:"mtime"  ` //更新时间
	Id     int64     `json:"id"  `
	Name   string    `json:"name"  `
	Nick   string    `json:"name"  `
	Pwd    string    `json:"pwd" `
	Phone  string    `json:"phone"`
	Email  string    `json:"email`
	Status int       `json:"status"  `
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"  `
	Did    string    `json:"did"  ` //冗余字段，为了查询方便
	Aid    string    `json:"adi"  `
	Openid string    `json: "openid"` //微信小程序唯一标识
	data   string    `json: "data"`   //备用字段

}

//需要其他的查询参数继续添加
type UserParams struct {
	Id     int64  `json:"id"`     //根据id查询
	Name   string `json:"name"`   //根据name模糊查询
	Status int    `json:"status"` //根据状态精确查询
	Phone  string `json:"phone"`  //根据手机号码模糊查询
	Did    string `json:"did"`    //根据域id精确查询
	Aid    string `json:"aid"`    //根据应用id精确查询
	Pn     int    `json:"pn"`
	Ps     int    `json:"ps"`
}

//需要其他的返回字段可以继续添加
type UserResult struct {
	Id     int64  `json:"id"`     //根据id查询
	Name   string `json:"name"`   //根据name模糊查询
	Status int    `json:"status"` //根据状态精确查询
	Phone  string `json:"phone"`  //根据手机号码模糊查询
	Did    string `json:"did"`    //根据域id精确查询
	Aid    string `json:"aid"`    //根据应用id精确查询
}
