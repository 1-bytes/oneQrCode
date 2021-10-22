package requests

// ValidateUserRegister 用户注册表单验证.
type ValidateUserRegister struct {
	Email    string `form:"email" binding:"required,email" label:"邮箱"`
	Username string `form:"username" binding:"required,min=6,max=20" label:"用户名"`
	Password string `form:"password" binding:"required,min=8,max=20" label:"密码"`
	Disable  bool   `form:"disable" binding:"-"`

	PasswordConfirm string `form:"password_confirm" binding:"required,min=8,max=20,eqfield=Password" label:"确认密码"`
	Captcha         string `form:"captcha" binding:"required,len=6" label:"验证码"`
}
