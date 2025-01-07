package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null;"`
	Email    string `json:"email" gorm:"unique; not null;"`
	Password string `json:"password" gorm:"not null;"`
	Role     string `json:"role" gorm:"default:user"`
}
