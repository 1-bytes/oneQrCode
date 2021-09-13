package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"oneQrCode/app/models/user"
	"oneQrCode/app/requests"
	"oneQrCode/pkg/captcha"
	"oneQrCode/pkg/config"
	"oneQrCode/pkg/logger"
)

type AuthController struct {
	BaseController
}

// DoRegister 用户注册.
func (ac *AuthController) DoRegister(c *gin.Context) {
	_user := user.User{
		Email:           c.PostForm("email"),
		Username:        c.PostForm("username"),
		Password:        c.PostForm("password"),
		PasswordConfirm: c.PostForm("password_confirm"),
		VerifyCode:      c.PostForm("verify_code"),
	}
	errs := requests.ValidateRegistrationForm(_user)
	if len(errs) > 0 {
		errCode := ac.getRandomString(16)

		c.JSON(500, gin.H{
			"message": "表单验证不通过，请验证修改后重新提交。",
			"error":   errs,
			"errCode": errCode,
		})
	}
}

// GetCaptcha 获取验证码.
func (ac *AuthController) GetCaptcha(c *gin.Context) {
	session := sessions.Default(c)
	decoder := json.NewDecoder(bytes.NewBufferString(config.GetString("captcha.style_json")))
	var verificationCode captcha.Captcha
	err := decoder.Decode(&verificationCode)
	logger.CheckError(err)
	id, b64s, err := verificationCode.NewCaptcha()
	logger.CheckError(err)
	session.Set("captcha_id", id)
	err = session.Save()
	logger.CheckError(err)
	c.JSON(200, gin.H{
		"result": b64s,
	})
}
