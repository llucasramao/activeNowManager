package router

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/activeNowManager/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/version", handlers.Version)
	}
}
