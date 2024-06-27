package main

import (
	"log"
	"sports-backend/infrastructure/db"
	"sports-backend/infrastructure/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()

	router.Run()
}
