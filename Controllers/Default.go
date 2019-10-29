package Controllers

import "github.com/gin-gonic/gin"


func DefaultResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}