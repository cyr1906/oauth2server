package main

import (
	"github.com/gin-gonic/gin"
	. "oauth2server/Controllers"
)

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()
	router := gin.Default()
	router.GET("/ping", DefaultResponse)
	router.GET("/authorize", Authorize)
	router.Run(":80") // 在 0.0.0.0:8080 上监听并服务
}
