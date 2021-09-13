package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"oneQrCode/app/http/controllers"
	"oneQrCode/pkg/config"
)

// RegisterMiddleware used for register middleware.
func RegisterMiddleware(r *gin.Engine) {
	store := cookie.NewStore([]byte(config.GetString("app.key")))
	r.Use(
		gin.Logger(),   // log
		gin.Recovery(), // exception handling
		sessions.Sessions("SESSIONS", store),
	)
}

// RegisterWebRoutes used for register router.
func RegisterWebRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		auth := controllers.AuthController{}
		authGroup.POST("/doRegister", auth.DoRegister)
		authGroup.GET("/getCaptcha", auth.GetCaptcha)
	}
}
