package bootstrap

import (
	"github.com/gin-gonic/gin"
	"oneQrCode/routes"
)

// SetupRoute used for init Router and middleware.
func SetupRoute() *gin.Engine {
	router := gin.New()
	routes.RegisterMiddleware(router)
	routes.RegisterWebRoutes(router)
	return router
}
