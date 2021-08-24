package bootstrap

import (
	"github.com/gin-gonic/gin"
	"oneQrCode/router"
)

// SetupRoute used for init Router and middleware.
func SetupRoute(r *gin.Engine) {
	router.RegisterMiddleware(r)
	router.RegisterWebRoutes(r)
}
