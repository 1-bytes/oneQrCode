package main

import (
	"github.com/gin-gonic/gin"
	"log"
	configs "oneQrCode/config"
	"oneQrCode/pkg/config"
)

func init() {
	configs.Initialize()
}

// main .
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":" + config.GetString("app.port"))
	if err != nil {
		log.Fatal("Run Service failed")
	}
}
