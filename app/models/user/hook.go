package user

import (
	"gorm.io/gorm"
	"oneQrCode/pkg/password"
)

// BeforeSave GORM 的钩子，在保存和更新模型之前调用.
func (user *User) BeforeSave(_ *gorm.DB) error {
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return nil
}
