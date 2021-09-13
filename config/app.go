package config

import "oneQrCode/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "oneQrCode"),
		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "develop"),
		// sessions 在 Cookie 中加密数据时使用
		"key": config.Env("APP_KEY", "H9ZBIEO2VY0R7XMQGWCKF43NDJTA8SUL15P6D5199EMIYGC4H4OALK0ZUJWD44PL"),
		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),
	})
}
