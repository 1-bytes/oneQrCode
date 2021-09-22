package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"oneQrCode/app/models/user"
	"oneQrCode/pkg/app"
	"oneQrCode/pkg/captcha"
	"oneQrCode/pkg/config"
	"oneQrCode/pkg/e"
	"oneQrCode/pkg/validation"
)

type AuthController struct {
	BaseController
}

// DoRegister 用户注册.
func (ac *AuthController) DoRegister(c *gin.Context) {
	appG := app.Gin{C: c}
	if err := c.ShouldBind(&user.User{}); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, validation.Translate(err))
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// GetCaptcha 获取验证码.
func (ac *AuthController) GetCaptcha(c *gin.Context) {
	appG := app.Gin{C: c}
	session := sessions.Default(c)
	decoder := json.NewDecoder(bytes.NewBufferString(config.GetString("captcha.style_json")))
	var verificationCode captcha.Captcha
	err := decoder.Decode(&verificationCode)
	if e.HasError(err) {
		appG.Response(http.StatusOK, e.ErrorGetCaptchaConfigFail, nil)
		return
	}
	id, b64s, err := verificationCode.NewCaptcha()
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
