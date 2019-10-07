package admin

/**
user session 用户session时长
*/

type UserSession struct {
	Uid     int64 `json:"uid"`
	Web     int   `json:"web"`
	Android int   `json:"android"`
	Ios     int   `json:"ios"`
}
