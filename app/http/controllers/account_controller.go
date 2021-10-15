package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"oneQrCode/app/models/user"
	"oneQrCode/pkg/app"
	"oneQrCode/pkg/captcha"
	"oneQrCode/pkg/e"
	"oneQrCode/pkg/validation"
)

type AccountController struct{}

// DoRegister 用户注册.
func (ac *AccountController) DoRegister(c *gin.Context) {
	appG := app.Gin{C: c}
	session := sessions.Default(c)
	// 表单验证
	var u user.User
	if err := c.ShouldBind(&u); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, validation.Translate(err))
		return
	}

	// 检查验证码合法性 无论成功与否，验证完成以后验证码失效
	captchaId := session.Get("captcha_id").(string)
	if !captcha.GetInstance().Verify(captchaId, u.Captcha, true) {
		appG.Response(http.StatusOK, e.ErrorVerifyCaptchaFail, nil)
		return
	}

	// 检查用户昵称和邮箱是否重复
	if user.HasUserByUsername(u.Username) || user.HasUserByEmail(u.Email) {
		appG.Response(http.StatusOK, e.ErrorExistUsernameOrEmail, nil)
		return
	}
	// 注册账号
	if err := u.Create(); err != nil {
		appG.Response(http.StatusOK, e.ErrorRegisterUserFail, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// GetCaptcha 获取验证码.
func (ac *AccountController) GetCaptcha(c *gin.Context) {
	appG := app.Gin{C: c}
	session := sessions.Default(c)
	id, b64s, err := captcha.GetInstance().Generate()
	if e.HasError(err) {
		appG.Response(http.StatusOK, e.ErrorInitCaptchaFail, nil)
		return
	}
	session.Set("captcha_id", id)
	err = session.Save()
	if e.HasError(err) {
		appG.Response(http.StatusOK, e.ErrorSaveSessionFail, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, b64s)
}
