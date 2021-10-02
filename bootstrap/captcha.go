package bootstrap

import (
	"oneQrCode/pkg/captcha"
	"oneQrCode/pkg/config"
)

// SetupCaptcha 初始化验证码.
func SetupCaptcha() {
	cfg := config.GetString("captcha.style_json")
	captcha.GetInstance(cfg)
}
