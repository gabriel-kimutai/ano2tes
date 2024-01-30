package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
    UserName string  `json:"username"`
	Email    string `gorm:"unique"`
    Password string	`json:"password"`
}