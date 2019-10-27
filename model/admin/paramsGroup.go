package admin

type ParamsGroup struct {
	Did  int64  `json:"did" binding:"required"`
	Aid  int64  `json:"aid" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ParamsUserGroup struct {
}
