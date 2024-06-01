package router

import (
	"net/http"
	"os"
	"sports-backend/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := SetRouter()

	router.Run(":8080")
}

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONT_URL"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	return router
}
