package controllers

import "github.com/gin-gonic/gin"

type AbcController struct {
}

func (a *AbcController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
