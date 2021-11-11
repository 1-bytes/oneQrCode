package requests

// ValidateUserLogin 用户登录表单验证.
type ValidateUserLogin struct {
	Email    string `form:"email" binding:"required,email" label:"邮箱"`
	Username string `form:"username" binding:"-"`
	Password string `form:"password" binding:"required,min=8,max=20" label:"密码"`
	Disable  bool   `form:"disable" binding:"-"`

	PasswordConfirm string `form:"password_confirm" binding:"-"`
	Captcha         string `form:"captcha" binding:"required,len=6" label:"验证码"`
}
