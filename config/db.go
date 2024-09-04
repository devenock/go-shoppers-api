package config

import (
	"github.com/Trend20/go-shopper-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("master.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//	migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	DB = db
}
