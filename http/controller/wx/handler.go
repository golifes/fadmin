package wx

import (
	"fadmin/logic/wx"
)

type HttpWxHandler struct {
	logic wx.LogicHandler
}

func NewHttpWxHandler(path string) *HttpWxHandler {
	return &HttpWxHandler{logic: wx.NewLogic(path)}
}
