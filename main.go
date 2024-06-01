package main

import (
	"sports-backend/db"
	"sports-backend/router"
)

func main() {
	db.Init()

	router.Run()
}
