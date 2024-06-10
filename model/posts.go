package model

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Date      time.Time      `json:"date"`
	Time      int            `json:"time"`
	Good      int            `json:"good"`
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
	Good      int            `json:"good"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"created_at"`
}
