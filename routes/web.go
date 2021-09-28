package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"oneQrCode/app/http/controllers"
	"oneQrCode/app/http/middlewares"
	"oneQrCode/pkg/config"
)

// RegisterMiddleware used for register middleware.
func RegisterMiddleware(r *gin.Engine) {
	store := cookie.NewStore([]byte(config.GetString("app.key")))
	r.Use(
		gin.Recovery(),                       // exception handling
		sessions.Sessions("SESSIONS", store), // SESSION
		middlewares.Logger(),                 // log
	)
}

// RegisterWebRoutes used for register router.
func RegisterWebRoutes(r *gin.Engine) {
	v := r.Group("api/v1")
	accountGroup := v.Group("account")
	{
		auth := controllers.AuthController{}
		accountGroup.POST("doRegister", auth.DoRegister)
		accountGroup.GET("getCaptcha", auth.GetCaptcha)
	}
}
