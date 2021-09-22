package user

import "oneQrCode/app/models"

// User 用户信息 Model.
type User struct {
	models.IdModel
	Email    string `gorm:"type:varchar(100);not null" valid:"email" form:"email" binding:"required,email" label:"邮箱"`
	Username string `gorm:"type:varchar(255);not null;unique" valid:"username" form:"username" binding:"required,min=6,max=20" label:"用户名"`
	Password string `gorm:"type:varchar(255);not null" valid:"password" form:"password" binding:"required,min=8,max=20" label:"密码"`
	Disable  bool   `gorm:"type:bool;not null;default:false" valid:"disable" form:"disable" binding:"-"`

	models.DateModel
	// gorm:"-" 设置 GORM 在读写时略过该字段
	PasswordConfirm string `gorm:"-" valid:"password_confirm" form:"password_confirm" binding:"required,min=8,max=20,eqfield=Password" label:"确认密码"`
	Captcha         string `gorm:"-" valid:"captcha" form:"captcha" binding:"required,len=6" label:"验证码"`
}
