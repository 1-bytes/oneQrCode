package bootstrap

import configs "oneQrCode/config"

// Setup 初始化指定的服务.
func Setup() {
	configs.Initialize()
	autoLoader(
		SetupDB,         // 数据库
		SetupLogs,       // 日志
		SetupValidation, // 表单验证
	)
}

func autoLoader(funcName ...func()) {
	// 只是单纯的初始化服务模块，没有参数，没有返回值！！
	for _, v := range funcName {
		v()
	}
}
