package config

import "oneQrCode/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "oneQrCode"),
		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "develop"),
		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),
		// 应用服务端口
		"port": config.Env("APP_PORT", "8080"),
		// gorilla/sessions 在 Cookie 中加密数据时使用
		"key": config.Env("APP_KEY", "YXOUi8AVM9r7uw3lGeoDQpdLSN6KgF1Hs2jC7bZE"),
		// 用来生成链接
		"url": config.Env("APP_URL", "http://127.0.0.1:3000"),
	})
}
