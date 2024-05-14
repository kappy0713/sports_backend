package main

import (
	"net/http"

	"sports-backend/migrate"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	migrate.Migrate()

	router.Run(":8080")
}
