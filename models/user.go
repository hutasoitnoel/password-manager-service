package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;notNull" json:"username" validate:"required"`
	Password string `gorm:"notNull" json:"password" validate:"required"`
}
