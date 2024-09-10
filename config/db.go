package config

import (
	"fmt"
	"github.com/Trend20/go-shoppers-api/models"
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
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		return
	}
	fmt.Println("Database connected successfully!")
	DB = db

}
