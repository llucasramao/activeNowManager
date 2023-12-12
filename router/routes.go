package router

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/activeNowManager/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/version", handlers.Version)
	router.POST("/NewReceived", handlers.NewReceived)
	router.GET("/getCommands", handlers.Commander)
	router.GET("/ViewReceiveds", handlers.ViewReciveds)
	router.GET("/linuxClient", handlers.InstallClientLinux)
	router.GET("/linuxClientBin", handlers.InstallClientLinuxBIN)
}
