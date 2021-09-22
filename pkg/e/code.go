package e

const (
	SUCCESS       = 200 // 成功
	ERROR         = 500 // 错误
	InvalidParams = 400 // 参数错误
	// -----------------
	ErrorSaveSessionFail      = 10001 // 存储 SESSION 失败
	ErrorGetCaptchaConfigFail = 10002 // 获取验证码参数失败
	ErrorInitCaptchaFail      = 10003 // 初始化验证码模块失败
	ErrorRegisterUserFail     = 10004 // 注册账号失败
	ErrorExistUsernameOrEmail = 10005 // 用户账号或邮件已经注册过了
)
