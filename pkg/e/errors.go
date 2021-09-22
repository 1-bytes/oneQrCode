package e

import (
	"oneQrCode/pkg/logger"
)

var MsgFlags = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "fail",
	InvalidParams:             "请求参数错误",
	ErrorSaveSessionFail:      "存储SESSION失败",
	ErrorGetCaptchaConfigFail: "获取验证码配置失败",
	ErrorInitCaptchaFail:      "初始化验证码模块失败",
}

// GetMsg get error information based on Code.
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// HasError any errors will be saved to the log.
func HasError(err error) bool {
	if err != nil {
		logger.Error(err)
		return true
	}
	return false
}
