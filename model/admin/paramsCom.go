package admin

type ParamsIds struct {
	Id     int64 `json:"id" binding:"required"`
	Status int   `json:"status"  binding:"required"`
}
