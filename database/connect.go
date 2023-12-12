package database

import (
	"fmt"
	"log"
	"os"

	"github.com/llucasramao/activeNowManager/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp" + "(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("Error connecting to database : error=%v", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&models.Received{}, &models.Port{})
	if err != nil {
		fmt.Println("Erro ao executar automigração:", err)
		os.Exit(1)
	}

	// reciber1 := models.Receiver{
	// 	Ip:       "192.168.1.1",
	// 	Hostname: "example1",
	// 	Os:       "Linux",
	// 	Ports: []models.Port{
	// 		{Port: 80},
	// 		{Port: 443},
	// 	},
	// }

	// reciber2 := models.Receiver{
	// 	Ip:       "192.168.1.2",
	// 	Hostname: "example2",
	// 	Os:       "Linux",
	// 	Ports: []models.Port{
	// 		{Port: 22},
	// 		{Port: 8080},
	// 	},
	// }

	// // Salvar o Reciber no banco de dados
	// db.Create(&reciber1)
	// db.Create(&reciber2)

	// var recibers []models.Receiver
	// db.Preload("Ports").Find(&recibers)

	// // Exibir os resultados
	// for _, reciber := range recibers {
	// 	fmt.Printf("Reciber ID: %d, IP: %s, Hostname: %s, OS: %s\n", reciber.ID, reciber.Ip, reciber.Hostname, reciber.Os)
	// 	for _, port := range reciber.Ports {
	// 		fmt.Printf("  Port ID: %d, Port Number: %d\n", port.ID, port.Port)
	// 	}
	// }

	DB = db
	return DB
}
