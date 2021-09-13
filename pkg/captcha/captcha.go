package captcha

import (
	"github.com/mojocn/base64Captcha"
)

// Captcha 验证码
type Captcha struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// NewCaptcha 创建一个验证码.
func (c *Captcha) NewCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	switch c.CaptchaType {
	case "audio":
		driver = c.DriverAudio
	case "string":
		driver = c.DriverString.ConvertFonts()
	case "math":
		driver = c.DriverMath.ConvertFonts()
	case "chinese":
		driver = c.DriverChinese.ConvertFonts()
	default:
		driver = c.DriverDigit
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	return id, b64s, err
}
