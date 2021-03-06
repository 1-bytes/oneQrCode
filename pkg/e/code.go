package e

const (
	SUCCESS       = 200 // 成功
	ERROR         = 500 // 错误
	InvalidParams = 400 // 参数错误
	// -----------------
	ErrorSaveSessionFail       = 10001 // 存储 SESSION 失败
	ErrorVerifyCaptchaFail     = 10002 // 校验验证码不通过
	ErrorInitCaptchaFail       = 10003 // 初始化验证码模块失败
	ErrorRegisterUserFail      = 10004 // 注册账号失败
	ErrorExistEmail            = 10005 // 该邮箱已经注册过了，请更换后重试
	ErrorLoginFail             = 10006 // 登录失败，账号或密码错误
	ErrorLoginDisabled         = 10007 // 登录失败，该账户已被封禁
	ErrorAuthCheckTokenFail    = 20001 // TOKEN 鉴权失败
	ErrorAuthCheckTokenOverdue = 20002 // TOKEN 已过期
	ErrorAuthToken             = 20003 // TOKEN 生成失败
	ErrorAuth                  = 20004 // TOKEN 错误
)
