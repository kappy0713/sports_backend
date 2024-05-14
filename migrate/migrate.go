package migrate

import (
	"fmt"
	"sports-backend/db"
	"sports-backend/model"
)

func Migrate() {
	fmt.Println("Migrating...")
	dbConnection := db.NewDB()
	defer fmt.Println("Successfully Migrated!!!!!")
	defer db.CloseDB(dbConnection)
	dbConnection.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Relationships{},
	)
}
