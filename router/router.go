package router

import (
	"os"
	"sports-backend/api"
	"sports-backend/middleware"

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

	router.POST("/", middleware.Auth(), api.Home)
	// ユーザー認証(新規登録・ログイン)
	router.POST("/register", api.Register)
	router.POST("/login", api.Login)

	// 運動記録投稿
	router.POST("/post", middleware.Auth(), api.Post)

	// 運動情報投稿
	router.POST("/share_post", middleware.Auth(), api.SharePost)

	return router
}
