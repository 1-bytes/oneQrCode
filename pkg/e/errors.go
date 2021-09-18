package e

import (
	"log"
)

var MsgFlags = map[int]string{
	SUCCESS:                      "ok",
	ERROR:                        "fail",
	INVALID_PARAMS:               "请求参数错误",
	ERROR_GET_CAPTCHACONFIG_FAIL: "获取验证码配置失败",
}

// GetMsg get error information based on Code.
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// CheckError any errors will be saved to the log.
func CheckError(err error) {
	if err != nil {
		log.Printf("[Error]: %s", err)
	}
}
