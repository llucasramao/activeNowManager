package router

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/activeNowManager/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/version", handlers.Version)
	router.POST("/receiver", handlers.Receiver)
	router.GET("/getCommands", handlers.Commander)
}
