package user

import (
	"oneQrCode/pkg/model"
	"strconv"
)

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

// HasUserByEmail 通过 Email 判断用户是否存在，存在返回 true，不存在返回 false.
func HasUserByEmail(email string) bool {
	var user User
	var count int64
	model.DB.Where("email = ?", email).First(&user).Count(&count)
	return count != 0
}

// HasUserByUsername 通过 Username 判断用户是否存在，存在返回 true，不存在返回 false.
func HasUserByUsername(username string) bool {
	var user User
	var count int64
	model.DB.Where("username = ?", username).First(&user).Count(&count)
	return count != 0
}

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
