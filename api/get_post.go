package api

import (
	"sports-backend/db"
	"sports-backend/model"

	"github.com/gin-gonic/gin"
)

// 運動記録の投稿を取得
func GetPost(c *gin.Context) {
	var posts []model.Post

	if err := db.DB.Order("created_at desc").Limit(10).Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	c.JSON(200, posts)
}

// 運動情報の投稿を取得
func GetSharePost(c *gin.Context) {
	var share_posts []model.SharePost

	if err := db.DB.Order("created_at desc").Limit(10).Find(&share_posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	c.JSON(200, share_posts)
}
