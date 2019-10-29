package Controllers

import "github.com/gin-gonic/gin"


func Authorize(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}