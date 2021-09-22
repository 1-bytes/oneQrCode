package app

import "C"
import (
	"github.com/gin-gonic/gin"
	"oneQrCode/pkg/e"
)

type Gin struct {
	C *gin.Context
}

// Response 返回的数据格式.
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response 设置 gin.JSON.
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}
