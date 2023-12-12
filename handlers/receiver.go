package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llucasramao/activeNowManager/database"
	"github.com/llucasramao/activeNowManager/models"
	"gorm.io/gorm"
)

func findOrCreatePort(db *gorm.DB, portNumber int) (*models.Port, error) {
	var existingPort models.Port

	result := db.Where("port = ?", portNumber).First(&existingPort)
	if result.Error == nil {
		return &existingPort, nil
	}

	newPort := models.Port{Port: portNumber}
	if err := db.Create(&newPort).Error; err != nil {
		return nil, err
	}

	return &newPort, nil
}
func NewReceived(c *gin.Context) {
	var received models.Received

	if err := c.ShouldBindJSON(&received); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se o IP já existe
	var existingReceived models.Received
	result := database.DB.Preload("Ports").Where("ip = ?", received.Ip).First(&existingReceived)

	if result.Error == nil {
		// O IP já existe, atualizar o registro existente
		existingReceived.Hostname = received.Hostname
		existingReceived.Os = received.Os

		// Excluir todas as portas associadas ao Received
		if err := database.DB.Model(&existingReceived).Association("Ports").Clear(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir as portas associadas"})
			return
		}

		// Verificar e criar/atualizar as portas associadas
		var receivedPorts []models.Port

		for _, portNumber := range received.Ports {
			port, err := findOrCreatePort(database.DB, portNumber.Port)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao encontrar/criar a porta"})
				return
			}
			receivedPorts = append(receivedPorts, *port)
		}

		// Atualizar a estrutura Received com as portas associadas
		existingReceived.Ports = receivedPorts

		// Salvar as alterações no banco de dados
		if err := database.DB.Save(&existingReceived).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o Received"})
			return
		}
		c.JSON(http.StatusOK, existingReceived)
	} else {
		// O IP não existe, criar um novo registro
		// (Note que as portas associadas também serão criadas no processo)
		if err := database.DB.Create(&received).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o Received"})
			return
		}
		c.JSON(http.StatusOK, received)
	}
}

// func NewReceived(c *gin.Context) {
// 	var received models.Received

// 	if err := c.ShouldBindJSON(&received); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var receivedPorts []models.Port

// 	for _, portNumber := range received.Ports {
// 		port, err := findOrCreatePort(database.DB, portNumber.Port)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao encontrar/criar a porta"})
// 			return
// 		}
// 		receivedPorts = append(receivedPorts, *port)
// 	}

// 	received.Ports = receivedPorts

// 	if err := database.DB.Create(&received).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o Received"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, received)
// }

func ViewReciveds(c *gin.Context) {
	var recibers []models.Received
	database.DB.Preload("Ports").Find(&recibers)
	c.JSON(200, gin.H{
		"status":  "geted",
		"content": recibers,
	})
}
