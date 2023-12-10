package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/llucasramao/activeNowManager/database"
	"github.com/llucasramao/activeNowManager/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	database.ConnectDB()
	router.Initialize()
}
