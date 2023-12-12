package handlers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(200, gin.H{"version": os.Getenv("VERSION")})
	c.Abort()
}

func Commander(c *gin.Context) {
	c.JSON(200, gin.H{"command": "null"})
	c.Abort()
}

func InstallClientLinux(c *gin.Context) {
	scriptPath := "./files/installClient.sh"

	// Lê o conteúdo do script Bash
	scriptContent, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler o script"})
		return
	}

	c.String(http.StatusOK, string(scriptContent))
}

func InstallClientLinuxBIN(c *gin.Context) {
	fileName := "main" // Nome do seu executável
	filePath := "./files/" + fileName

	c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filePath)
}
