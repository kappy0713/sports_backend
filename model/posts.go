package model

import (
	"time"

	"github.com/lib/pq"
)

// 運動記録
type Post struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Date      time.Time      `json:"date"`
	Time      int            `json:"time"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"created_at"`
}

type PostResponse struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Date      time.Time      `json:"date"`
	Time      int            `json:"time"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"created_at"`
}

// 運動情報
type SharePost struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	URL       string         `json:"url"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"created_at"`
}

type SharePostResponse struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	URL       string         `json:"url"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"created_at"`
}
