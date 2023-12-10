package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(200, gin.H{"version": os.Getenv("VERSION")})
	c.Abort()
}
