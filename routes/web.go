package routes

import (
	"github.com/gin-gonic/gin"
	"oneQrCode/app/controllers"
)

// RegisterMiddleware used for register middleware.
func RegisterMiddleware(r *gin.Engine) {
	r.Use(
		gin.Logger(),   // log
		gin.Recovery(), // exception handling
	)
}

// RegisterWebRoutes used for register router.
func RegisterWebRoutes(r *gin.Engine) {
	abc := controllers.AbcController{}
	r.GET("/ping", abc.Index)
}
