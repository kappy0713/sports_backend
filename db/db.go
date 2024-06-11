package db

import (
	"fmt"
	"log"
	"os"
	"sports-backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	DB = db

	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.SharePost{}, &model.Relationships{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Succesfully Connected!!")
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
