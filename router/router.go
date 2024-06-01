package router

import (
	"net/http"
	"sports-backend/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := SetRouter()

	router.Run(":8080")
}

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	return router
}
