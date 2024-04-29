package models

import "gorm.io/gorm"

type Credential struct {
	gorm.Model
	UserId      uint   `gorm:"notNull" json:"user_id"`
	WebsiteUrl  string `gorm:"notNull" json:"website_url" validate:"required"`
	WebsiteName string `gorm:"notNull" json:"website_name" validate:"required"`
	Username    string `gorm:"notNull" json:"username" validate:"required"`
	Password    string `gorm:"notNull" json:"password" validate:"required"`
}
