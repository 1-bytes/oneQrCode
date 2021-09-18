package app

import (
	"github.com/gin-gonic/gin"
	"oneQrCode/pkg/e"
	"oneQrCode/pkg/logger"
	"oneQrCode/pkg/utils"
)

type Gin struct {
	c *gin.Context
}

// Response 返回的数据格式.
type Response struct {
	UniqueID string      `json:"unique_id"`
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
}

// Response 设置 gin.JSON.
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	// 为错误标记个唯一性质的ID返回
	uniqueID := ""
	var err error
	if httpCode != e.SUCCESS {
		uniqueID, err = utils.GetUUID()
		if err != nil {
			uniqueID = utils.GetRandomString(32)
		}
		logger.Error()
	}
	g.c.JSON(httpCode, Response{
		UniqueID: uniqueID,
		Code:     errCode,
		Msg:      e.GetMsg(errCode),
		Data:     data,
	})
}
