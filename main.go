package main

import (
	"log"
	"sports-backend/db"
	"sports-backend/router"

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
