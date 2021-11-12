package routes

import (
	"github.com/gin-gonic/gin"
	"oneQrCode/app/http/controllers"
)

// Register used for register router.
func Register(r *gin.Engine) {
	v := r.Group("api/v1")
	userGroup := v.Group("user")
	{
		controller := controllers.UserController{}
		userGroup.GET("getCaptcha", controller.GetCaptcha)
		userGroup.POST("doRegister", controller.DoRegister)
		userGroup.POST("doLogin", controller.DoLogin)
	}
}
