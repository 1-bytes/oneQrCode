package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"oneQrCode/pkg/config"
	"oneQrCode/pkg/logger"
	"time"
)

// Register used for register middleware.
func Register(r *gin.Engine) {
	store := cookie.NewStore([]byte(config.GetString("app.key")))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   config.GetString("http.listen_host"),
		MaxAge:   86400 * 3,
		Secure:   false,
		HttpOnly: true,
		SameSite: 0,
	})

	r.Use(
		RequestId(),                               // 为每个请求标记一个唯一性质的 ID
		GinRecoveryWithZap(logger.Logger, true),   // err 和 panic 记录到日志（包括堆栈信息）
		GinZap(logger.Logger, time.RFC3339, true), // 访问请求记录到日志
		sessions.Sessions("SESSIONS", store),      // SESSION
	)
}
