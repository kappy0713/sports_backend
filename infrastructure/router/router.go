package router

import (
	"os"
	"sports-backend/infrastructure/middleware"
	"sports-backend/interfaces/controller"

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

	// router.POST("/", middleware.Auth(), controller.Home)
	// ユーザー認証(新規登録・ログイン)
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	// 運動記録投稿
	router.POST("/post", middleware.Auth(), controller.Post)

	// 運動情報投稿
	router.POST("/share_post", middleware.Auth(), controller.SharePost)

	// 運動記録一覧
	router.GET("/list", controller.GetPost)

	// 共有情報一覧
	router.GET("/share", controller.GetSharePost)

	// マイページ
	router.POST("/name", middleware.Auth(), controller.UserName)
	router.POST("/mypage", middleware.Auth(), controller.GetMyPost)
	router.POST("/time", middleware.Auth(), controller.GetTime)
	router.POST("/month", middleware.Auth(), controller.GetMonthTime)

	return router
}
