package bootstrap

import (
	"oneQrCode/app/requests"
	"oneQrCode/pkg/validation"
)

// SetupValidation 初始化表单验证器.
func SetupValidation() {
	validation.Init()
	requests.RegisterTranslation()
}
