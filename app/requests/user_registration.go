package requests

import (
	"encoding/json"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"oneQrCode/app/models/user"
	"oneQrCode/pkg/e"
	"oneQrCode/pkg/validation"
)

// UserRegistration 表单校验.
type UserRegistration struct {
	Email           string `validate:"required,email" label:"邮箱"`
	Username        string `validate:"required,min=6,max=20" label:"用户名"`
	Password        string `validate:"required,min=8,max=20" label:"密码"`
	PasswordConfirm string `validate:"required,min=8,max=20,eqfield=Password" label:"确认密码"`
}

// ValidateRegistrationForm 用户注册表单验证.
func ValidateRegistrationForm(data user.User) validation.Result {
	// 类型转换
	var u UserRegistration
	m, _ := json.Marshal(data)
	err := json.Unmarshal(m, &u)
	if err != nil {
		u = UserRegistration{}
	}

	// 单例初始化验证器
	trans, validate, err := validation.GetInstance().SetupValidator()
	e.CheckError(err)

	// 自定义翻译
	validate.RegisterTranslation("eqfield=Password", trans, func(ut ut.Translator) error { //nolint:errcheck
		return ut.Add("eqfield", "{0}失败，请检查两次输入的密码是否正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("eqfield", fe.Field())
		return t
	})

	// 表单验证
	err = validate.Struct(u)
	errs := validation.Result{}
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.StructField()
			errs[field] = append(errs[field], err.Translate(trans))
		}
	}
	return errs
}
