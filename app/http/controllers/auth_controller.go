package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"oneQrCode/app/models/user"
	"oneQrCode/app/requests"
	"oneQrCode/pkg/captcha"
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
	config := `{
				  "ShowLineOptions": [
					"2",
					"4",
					"8"
				  ],
				  "CaptchaType": "string",
				  "Id": "",
				  "VerifyValue": "duku3m",
				  "DriverAudio": {
					"Length": 6,
					"Language": "zh"
				  },
				  "DriverString": {
					"Height": 60,
					"Width": 230,
					"ShowLineOptions": 14,
					"NoiseCount": 480,
					"Source": "234568wertyupkjhfdsazxcvnm",
					"Length": 6,
					"Fonts": [
					  "Flim-Flam.ttf",
					  "DeborahFancyDress.ttf",
					  "actionj.ttf",
					  "RitaSmith.ttf",
					  "chromohv.ttf",
					  "ApothecaryFont.ttf"
					],
					"BgColor": {
					  "R": 0,
					  "G": 0,
					  "B": 0,
					  "A": 0
					}
				  },
				  "DriverMath": {
					"Height": 60,
					"Width": 240,
					"ShowLineOptions": 14,
					"NoiseCount": 0,
					"Length": 6,
					"Fonts": [
					  "wqy-microhei.ttc"
					],
					"BgColor": {
					  "R": 0,
					  "G": 0,
					  "B": 0,
					  "A": 0
					}
				  },
				  "DriverChinese": {
					"Height": 60,
					"Width": 320,
					"ShowLineOptions": 0,
					"NoiseCount": 0,
					"Source": "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值",
					"Length": 2,
					"Fonts": [
					  "wqy-microhei.ttc"
					],
					"BgColor": {
					  "R": 125,
					  "G": 125,
					  "B": 0,
					  "A": 118
					}
				  },
				  "DriverDigit": {
					"Height": 80,
					"Width": 240,
					"Length": 5,
					"MaxSkew": 0.7,
					"DotCount": 80
				  }
				}`
	session := sessions.Default(c)
	decoder := json.NewDecoder(bytes.NewBufferString(config))
	var verificationCode captcha.Captcha
	err := decoder.Decode(&verificationCode)
	logger.CheckError(err)
	id, b64s, err := verificationCode.NewCaptcha()
	logger.CheckError(err)
	session.Set("captcha_id", id)
	c.JSON(200, gin.H{
		"result": b64s,
	})
}
