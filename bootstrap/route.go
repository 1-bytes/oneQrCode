package bootstrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oneQrCode/app/http/middlewares"
	"oneQrCode/routes"
)

// SetupRoute used for init Router and middleware.
func SetupRoute() http.Handler {
	router := gin.New()
	middlewares.Register(router)
	routes.Register(router)
	return router
}
