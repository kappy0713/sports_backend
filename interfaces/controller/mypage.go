package controller

import (
	"net/http"
	"sports-backend/domain"
	"sports-backend/infrastructure/db"
	"time"

	"github.com/gin-gonic/gin"
)

type Time struct {
	Date string `json:"date"`
	Time int    `json:"time"`
}

// ユーザーネームを返す
func UserName(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	u := user.(domain.User)

	c.JSON(200, gin.H{"name": u.Name})
}

// 直近7日間の運動時間を取得
func GetTime(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	u := user.(domain.User)

	location, _ := time.LoadLocation("Asia/Tokyo")
	today := time.Now().In(location)
	eightdays := today.AddDate(0, 0, -7)

	id := u.ID
	var posts []domain.Post

	times := make([]Time, 7)
	for i := 0; i < 7; i++ {
		day := eightdays.AddDate(0, 0, i+1)
		times[i] = Time{
			Date: day.Format("2006-01-02"),
			Time: 0,
		}
	}

	// DBから直近7日間の投稿を取得
	if err := db.DB.Where("user_id = ? AND date BETWEEN ? AND ?", id, eightdays, today).Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	// 各日の運動時間の合計を計算
	for _, post := range posts {
		date := post.Date.Format("2006-01-02")
		for i := range times {
			if times[i].Date == date {
				times[i].Time += post.Time
				break
			}
		}
	}

	c.JSON(200, times)
}

// 今月の運動時間を取得
func GetMonthTime(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	u := user.(domain.User)
	id := u.ID

	location, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(location)
	first := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, location)

	monthtimes := []Time{}
	for d := first; d.Before(now) || d.Equal(now); d = d.AddDate(0, 0, 1) {
		monthtimes = append(monthtimes, Time{
			Date: d.Format("2006-01-02"),
			Time: 0,
		})
	}

	var posts []domain.Post
	// DBから今月の投稿を取得
	if err := db.DB.Where("user_id = ? AND date BETWEEN ? AND ?", id, first, now).Order("date ASC").Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	// 各日の運動時間の合計を計算
	for _, post := range posts {
		postDate := post.Date.Format("2006-01-02")
		for i := range monthtimes {
			if monthtimes[i].Date == postDate {
				monthtimes[i].Time += post.Time
				break
			}
		}
	}

	c.JSON(200, monthtimes)
}

// 自分の投稿を取得
func GetMyPost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	u := user.(domain.User)
	id := u.ID

	var posts []domain.Post
	// DBから自分の投稿を取得
	if err := db.DB.Where("user_id = ?", id).Order("created_at desc").Limit(10).Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	c.JSON(200, posts)
}
