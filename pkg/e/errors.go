package e

import "oneQrCode/pkg/logger"

var MsgFlags = map[int]string{
	SUCCESS:                    "success",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorSaveSessionFail:       "存储 SESSION 失败",
	ErrorVerifyCaptchaFail:     "校验验证码不通过",
	ErrorInitCaptchaFail:       "初始化验证码模块失败",
	ErrorRegisterUserFail:      "注册账号失败",
	ErrorExistEmail:            "该邮箱已经注册过了，请更换后重试",
	ErrorLoginFail:             "登录失败，账号或密码错误",
	ErrorLoginDisabled:         "登录失败，该账户已被封禁",
	ErrorAuthCheckTokenFail:    "TOKEN 鉴权失败",
	ErrorAuthCheckTokenOverdue: "TOKEN 已过期",
	ErrorAuthToken:             "TOKEN 生成失败",
	ErrorAuth:                  "TOKEN 错误",
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
		sugar := logger.Logger.Sugar()
		sugar.Errorf("An unpredictable error was caught: %s", err)
		return true
	}
	return false
}
