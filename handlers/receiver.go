package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llucasramao/activeNowManager/database"
	"github.com/llucasramao/activeNowManager/models"
	"gorm.io/gorm"
)

func findOrCreatePort(db *gorm.DB, portNumber int) (*models.Port, error) {
	var port models.Port
	if err := db.Where("port = ?", portNumber).FirstOrCreate(&port, models.Port{Port: portNumber}).Error; err != nil {
		return nil, err
	}
	return &port, nil
}

func findOrCreateApp(db *gorm.DB, appName, appVersion string) (*models.App, error) {
	var app models.App
	if err := db.Where("name = ? AND version = ?", appName, appVersion).FirstOrCreate(&app, models.App{Name: appName, Version: appVersion}).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func NewReceived(c *gin.Context) {
	var receivedData models.Received
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var received models.Received
	err := database.DB.Where("ip = ?", receivedData.Ip).Preload("Ports").Preload("Apps").FirstOrCreate(&received, models.Received{Ip: receivedData.Ip}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao encontrar/criar o Received"})
		return
	}

	// Atualizando hostname e OS
	received.Hostname = receivedData.Hostname
	received.Os = receivedData.Os

	// Associando Portas e Apps
	for _, portData := range receivedData.Ports {
		port, err := findOrCreatePort(database.DB, portData.Port)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao encontrar/criar a porta"})
			return
		}
		received.Ports = append(received.Ports, *port)
	}

	for _, appData := range receivedData.Apps {
		app, err := findOrCreateApp(database.DB, appData.Name, appData.Version)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao encontrar/criar o App"})
			return
		}
		received.Apps = append(received.Apps, *app)
	}

	if err := database.DB.Save(&received).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar o Received"})
		return
	}

	c.JSON(http.StatusOK, received)
}

// func NewReceived(c *gin.Context) {
// 	var received models.Received

// 	if err := c.ShouldBindJSON(&received); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var existingReceived models.Received
// 	result := database.DB.Preload("Ports").Where("ip = ?", received.Ip).First(&existingReceived)

// 	if result.Error == nil {
// 		existingReceived.Hostname = received.Hostname
// 		existingReceived.Os = received.Os

// 		if err := database.DB.Model(&existingReceived).Association("Ports").Clear(); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir as portas associadas"})
// 			return
// 		}

// 		var receivedPorts []models.Port

// 		for _, portNumber := range received.Ports {
// 			port, err := findOrCreatePort(database.DB, portNumber.Port)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao encontrar/criar a porta"})
// 				return
// 			}
// 			receivedPorts = append(receivedPorts, *port)
// 		}

// 		existingReceived.Ports = receivedPorts

// 		if err := database.DB.Save(&existingReceived).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o Received"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, existingReceived)
// 	} else {
// 		if err := database.DB.Create(&received).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o Received"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, received)
// 	}
// }

func ViewReciveds(c *gin.Context) {
	var recibers []models.Received
	database.DB.Preload("Ports").Preload("Apps").Find(&recibers)
	c.JSON(200, gin.H{
		"status":  "geted",
		"content": recibers,
	})
}
