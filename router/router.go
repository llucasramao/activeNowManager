package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Permitir solicitações de qualquer origem (isso pode ser ajustado conforme necessário)
	config.AddAllowHeaders("Authorization")

	router.Use(cors.New(config))

	InitializeRoutes(router)
	router.Run(":7654")
}
