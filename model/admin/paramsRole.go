package admin

type ParamsRole struct {
	Did  int64  `json:"did" binding:"required" ` //域id
	Aid  int64  `json:"aid"  binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ParamsRoleDel struct {
	Id int64 `json:"id" binding:"required"`
}

type ParamsRoleList struct {
	Id     int64  `json:"id"`
	Did    int64  `json:"did"`
	Aid    int64  `json:"aid"`
	Rid    int64  `json:"rid"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Ctime  int    `json:"ctime"`
	Mtime  int    `json:"mtime"`
	Ps     int    `json:"ps,omitempty"`
	Pn     int    `json:"pn,omitempty"`
}
