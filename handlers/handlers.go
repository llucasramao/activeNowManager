package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(200, gin.H{"version": os.Getenv("VERSION")})
	c.Abort()
}

func Receiver(c *gin.Context) {
	c.JSON(200, "recept")
	c.Abort()
}

func Commander(c *gin.Context) {
	c.JSON(200, gin.H{"command": "null"})
	c.Abort()
}
