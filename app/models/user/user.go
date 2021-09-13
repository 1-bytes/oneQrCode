package user

import "oneQrCode/app/models"

// User 用户信息 Model.
type User struct {
	models.IdModel
	Email    string `gorm:"type:varchar(100);not null" valid:"email"`
	Username string `gorm:"type:varchar(255);not null;unique" valid:"username"`
	Password string `gorm:"type:varchar(255);not null" valid:"password"`
	Disable  bool   `gorm:"type:bool;;not null;default:false" valid:"disable"`

	models.DateModel
	// gorm:"-" 设置 GORM 在读写时略过该字段
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
	VerifyCode      string `gorm:"-" valid:"verify_code"`
}
