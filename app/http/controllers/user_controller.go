package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"oneQrCode/app/models/user"
	"oneQrCode/app/requests"
	"oneQrCode/pkg/captcha"
	"oneQrCode/pkg/e"
	"oneQrCode/pkg/validation"
	"strconv"
)

type UserController struct {
	BaseController
}

// GetCaptcha 获取验证码.
func (uc *UserController) GetCaptcha(c *gin.Context) {
	appG := uc.GetAppG(c)
	session := uc.GetSessions(c)
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

// DoRegister 用户注册.
func (uc *UserController) DoRegister(c *gin.Context) {
	appG := uc.GetAppG(c)
	session := uc.GetSessions(c)
	// 表单验证
	var data requests.ValidateUserRegister
	if err := c.ShouldBind(&data); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, validation.Translate(err))
		return
	}
	// 检查验证码合法性 无论成功与否，验证完成以后验证码失效
	captchaId, _ := session.Get("captcha_id").(string)
	session.Delete("captcha_id")
	_ = session.Save()
	if !captcha.GetInstance().Verify(captchaId, data.Captcha, true) {
		appG.Response(http.StatusOK, e.ErrorVerifyCaptchaFail, nil)
		return
	}
	// 检查用户邮箱是否重复
	if user.HasByEmail(data.Email) {
		appG.Response(http.StatusOK, e.ErrorExistEmail, nil)
		return
	}
	// 注册账号
	var u user.User
	m, _ := json.Marshal(data)
	_ = json.Unmarshal(m, &u)
	if err := u.Create(); err != nil {
		appG.Response(http.StatusOK, e.ErrorRegisterUserFail, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// DoLogin 用户登录.
func (uc *UserController) DoLogin(c *gin.Context) {
	appG := uc.GetAppG(c)
	session := uc.GetSessions(c)
	// 表单验证
	var data requests.ValidateUserLogin
	if err := c.ShouldBind(&data); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, validation.Translate(err))
		return
	}
	// 检查验证码合法性 无论成功与否，验证完成以后验证码失效
	captchaId, _ := session.Get("captcha_id").(string)
	session.Delete("captcha_id")
	_ = session.Save()
	if !captcha.GetInstance().Verify(captchaId, data.Captcha, true) {
		appG.Response(http.StatusOK, e.ErrorVerifyCaptchaFail, nil)
		return
	}
	// 检查邮箱和对应的密码是否正确
	userInfo, err := user.GetByEmail(data.Email)
	if err != nil || !user.CheckPassword(data.Password, userInfo.Password) {
		appG.Response(http.StatusOK, e.ErrorLoginFail, nil)
		return
	}
	// 检查账户是否被封禁
	if userInfo.Disable {
		appG.Response(http.StatusOK, e.ErrorLoginDisabled, nil)
		return
	}
	info := map[string]string{
		"uid":      strconv.FormatUint(userInfo.ID, 10),
		"email":    userInfo.Email,
		"username": userInfo.Username,
	}
	session.Set("user_info", info)
	_ = session.Save()
	appG.Response(http.StatusOK, e.SUCCESS, info)
}

// Logout 退出登录.
func (uc UserController) Logout(c *gin.Context) {
	// TODO: 待开发
}
