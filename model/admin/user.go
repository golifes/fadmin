package admin

/**
加上冗余两个字段
*/
type User struct {
	//Id  int64
	Id     int64  `json:"id"  `                 //主键id
	Name   string `json:"name"  `               //用户名
	Nick   string `json:"name"  `               //昵称
	Pwd    string `json:"pwd" `                 //密码
	Phone  string `json:"phone"`                //电话号码
	Status int    `json:"status"`               //
	Ctime  int    `json:"ctime" xorm:"created"` //创建时间
	Mtime  int    `json:"mtime" xorm:"updated"` //更新时间
	Did    int64  `json:"did"`                  //
	Aid    int64  `json:"aid"`                  //
	Gid    int    `json:"gid"`                  //
	Openid string `json:"openid"`               //微信小程序唯一标识
	Data   string `json:"data"`                 //备用字段
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
