package models

import "gorm.io/gorm"

type Saving struct {
	gorm.Model
	UserId      uint   `gorm:"notNull" json:"user_id"`
	Name        string `gorm:"notNull" json:"name" validate:"required"`
	Amount      int    `gorm:"notNull" json:"amount" validate:"required"`
	Description string `gorm:"notNull" json:"description"`
}
