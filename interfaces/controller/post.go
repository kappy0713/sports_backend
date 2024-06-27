package controller

import (
	"net/http"
	"sports-backend/domain"
	"sports-backend/infrastructure/db"

	"github.com/gin-gonic/gin"
)

// 運動記録を投稿
func Post(c *gin.Context) {
	var post domain.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	userInfo, ok := user.(domain.User)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}
	post.UserID = userInfo.ID
	post.Name = userInfo.Name

	if err := db.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering user"})
		return
	}

	c.Status(http.StatusOK)
}

// 運動の情報を投稿
func SharePost(c *gin.Context) {
	var share_post domain.SharePost
	if err := c.BindJSON(&share_post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	userInfo, ok := user.(domain.User)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}
	share_post.UserID = userInfo.ID
	share_post.Name = userInfo.Name

	if err := db.DB.Create(&share_post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering user"})
		return
	}

	c.Status(http.StatusOK)
}
