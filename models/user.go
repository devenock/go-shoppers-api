package models

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Pass  string `json:"pass" binding:"required"`
	Role  string `json:"role" binding:"required"`
}
