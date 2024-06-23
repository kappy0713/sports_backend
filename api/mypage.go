package api

import (
	"sports-backend/db"
	"sports-backend/model"
	"time"

	"github.com/gin-gonic/gin"
)

type Time struct {
	Date string `json:"date"`
	Time int    `json:"time"`
}

// 直近7日間の運動時間を取得
func GetTime(c *gin.Context) {
	// user, exists := c.Get("user")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// u := user.(*model.User)

	location, _ := time.LoadLocation("Asia/Tokyo")
	today := time.Now().In(location)
	sevendays := today.AddDate(0, 0, -6)

	id := 1
	var posts []model.Post

	times := make([]Time, 7)
	for i := 0; i < 7; i++ {
		day := sevendays.AddDate(0, 0, i)
		times[i] = Time{
			Date: day.Format("2006-01-02"),
			Time: 0,
		}
	}

	// DBから直近7日間の投稿を取得
	if err := db.DB.Where("user_id = ? AND date BETWEEN ? AND ?", id, sevendays, today).Find(&posts).Error; err != nil {
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
	// user, exists := c.Get("user")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// u := user.(*model.User)
	id := 1

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

	var posts []model.Post
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
	// user, exists := c.Get("user")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// u := user.(model.User)
	id := 1

	var posts []model.Post
	// DBから自分の投稿を取得
	if err := db.DB.Where("user_id = ?", id).Order("created_at desc").Limit(10).Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	c.JSON(200, posts)
}
