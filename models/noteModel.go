package models

import "gorm.io/gorm"

type Note struct {
    gorm.Model
    Message string `json:"message"`
}
