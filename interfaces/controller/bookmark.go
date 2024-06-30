package controller

import (
	"fmt"
	"net/http"
	"sports-backend/domain"
	"sports-backend/infrastructure/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ブックマークした投稿をDBに登録する関数
func PostBookmark(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	u := user.(domain.User)

	var ID domain.BookmarkRequest

	if err := c.ShouldBindJSON(&ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 同一のIDが存在するか確認
	for _, id := range u.Bookmark {
		if id == strconv.Itoa(ID.ID) {
			c.JSON(http.StatusOK, gin.H{"message": "ID already exists in bookmarks"})
			return
		}
	}

	u.Bookmark = append(u.Bookmark, strconv.Itoa(ID.ID))

	if err := db.DB.Model(&u).Update("Bookmark", u.Bookmark).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating bookmarks"})
		return
	}

	c.JSON(200, gin.H{"message": "Bookmark added"})

}

// ブックマークした投稿を返す関数
func GetBookmark(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	u := user.(domain.User)

	IDs := u.Bookmark
	var intIDs []int
	var share_posts []domain.SharePost

	for _, id := range IDs {
		intID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error: ID is not an int type")
			continue
		}
		intIDs = append(intIDs, intID)
	}

	if err := db.DB.Where("id IN ?", intIDs).Order("created_at desc").Find(&share_posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(200, share_posts)
}
