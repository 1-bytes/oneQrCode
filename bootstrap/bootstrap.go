package bootstrap

// Setup 初始化指定的服务.
func Setup() {
	loader(
		SetupDB,   // 数据库
		SetupLogs, // 日志
	)
}

func loader(funcName ...func()) {
	// 只是单纯的初始化服务模块，没有参数，没有返回值！！
	for _, v := range funcName {
		v()
	}
}
