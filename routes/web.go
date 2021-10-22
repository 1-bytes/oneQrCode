package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"oneQrCode/app/http/controllers"
	"oneQrCode/app/http/middlewares"
	"oneQrCode/pkg/config"
	"oneQrCode/pkg/logger"
	"time"
)

// RegisterMiddleware used for register middleware.
func RegisterMiddleware(r *gin.Engine) {
	store := cookie.NewStore([]byte(config.GetString("app.key")))

	r.Use(
		middlewares.RequestId(),                               // 为每个请求标记一个唯一性质的 ID
		middlewares.GinRecoveryWithZap(logger.Logger, true),   // err 和 panic 记录到日志（包括堆栈信息）
		middlewares.GinZap(logger.Logger, time.RFC3339, true), // 访问请求记录到日志
		sessions.Sessions("SESSIONS", store),                  // SESSION
	)
}

// RegisterWebRoutes used for register router.
func RegisterWebRoutes(r *gin.Engine) {
	v := r.Group("api/v1")
	userGroup := v.Group("user")
	{
		controller := controllers.UserController{}
		userGroup.GET("getCaptcha", controller.GetCaptcha)
		userGroup.POST("doRegister", controller.DoRegister)
	}
}
