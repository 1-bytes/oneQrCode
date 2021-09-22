package requests

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"oneQrCode/pkg/validation"
)

func RegisterTranslation() {
	_ = validation.Validate.RegisterTranslation("eqfield=Password", validation.Trans, func(ut ut.Translator) error {
		return ut.Add("eqfield", "{0}失败，请检查两次输入的密码是否正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("eqfield", fe.Field())
		return t
	})
}
