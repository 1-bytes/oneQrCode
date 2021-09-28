package bootstrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oneQrCode/routes"
)

// SetupRoute used for init Router and middleware.
func SetupRoute() http.Handler {
	router := gin.New()
	routes.RegisterMiddleware(router)
	routes.RegisterWebRoutes(router)
	return router
}
