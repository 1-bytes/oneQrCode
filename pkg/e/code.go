package e

const (
	SUCCESS        = 200 // 成功
	ERROR          = 500 // 错误
	INVALID_PARAMS = 400 // 参数错误
	// -----------------
	ERROR_SAVE_SESSION_FAIL      = 10001 // 存储 SESSION 失败
	ERROR_GET_CAPTCHACONFIG_FAIL = 10002 // 获取验证码参数失败
	ERROR_INIT_CAPTCHA_FAIL      = 10003 // 初始化验证码模块失败
)
