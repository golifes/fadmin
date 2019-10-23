package admin

type ParamsRole struct {
	Did  int64  `json:"did" binding:"required" ` //域id
	Aid  int64  `json:"aid"  binding:"required"`
	Name string `json:"name" binding:"required"`
}
