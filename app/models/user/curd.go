package user

import (
	"oneQrCode/pkg/model"
	"oneQrCode/pkg/password"
	"strconv"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功.
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新用户资料.
func (user *User) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&user)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// Get 根据 ID 获取用户信息.
func Get(uidStr string) (User, error) {
	var user User
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return User{}, err
	}
	if err := model.DB.First(&user, uid).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// GetByEmail 根据 Email 获取用户信息.
func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// HasByEmail 通过 Email 判断用户是否存在，存在返回 true，不存在返回 false.
func HasByEmail(email string) bool {
	var user User
	var count int64
	model.DB.Where("email = ?", email).First(&user).Count(&count)
	return count != 0
}

// CheckPassword 校验密码是否正确，成功返回 true，失败返回false.
func CheckPassword(pass, hash string) bool {
	return password.IsHashed(hash) && password.CheckHash(pass, hash)
}
